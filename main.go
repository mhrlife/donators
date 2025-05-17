package main

import (
	"context"
	"log/slog"
	"os"
	"strconv"
	"time"

	"DonateNotifier/database"
	"DonateNotifier/ent"
	"DonateNotifier/reymit"
	"DonateNotifier/telegram"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file", "error", err)
		// Continue execution, as env vars might be set in the environment
	}

	slog.SetLogLoggerLevel(slog.LevelDebug)

	botToken := os.Getenv("BOT_TOKEN")
	groupIDStr := os.Getenv("GROUP_ID")
	reymitToken := os.Getenv("REYMIT_TOKEN")

	if botToken == "" {
		slog.Error("BOT_TOKEN environment variable not set")
		os.Exit(1)
	}
	if groupIDStr == "" {
		slog.Error("GROUP_ID environment variable not set")
		os.Exit(1)
	}
	if reymitToken == "" {
		slog.Error("REYMIT_TOKEN environment variable not set")
		os.Exit(1)
	}

	groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
	if err != nil {
		slog.Error("Invalid GROUP_ID environment variable", "error", err)
		os.Exit(1)
	}

	// Initialize the database client
	dbClient, err := database.InitDB()
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	defer dbClient.Close()
	slog.Info("Database client initialized")

	// Initialize the Telegram bot
	bot, err := telegram.NewBot(botToken, groupID)
	if err != nil {
		slog.Error("Failed to initialize Telegram bot", "error", err)
		os.Exit(1)
	}
	slog.Info("Telegram bot initialized")

	// Use a context for graceful shutdown later if needed
	ctx := context.Background()

	// Main loop to fetch and process donations periodically
	slog.Info("Starting donation fetch loop")
	for {
		slog.Debug("Fetching latest donations from Reymit API")
		resp, err := reymit.GetLastDonations(reymitToken)
		if err != nil {
			slog.Error("Failed to fetch donations from Reymit API", "error", err)
			// Wait before retrying
			time.Sleep(30 * time.Second)
			continue
		}

		if !resp.Ok {
			slog.Error("Reymit API returned non-ok response", "response", resp)
			// Wait before retrying
			time.Sleep(30 * time.Second)
			continue
		}

		slog.Debug("Received donations from Reymit API", "count", len(resp.Data.Donates))

		// Process donations in reverse order (newest first)
		// This helps in stopping early if we hit an already processed donation
		for i := len(resp.Data.Donates) - 1; i >= 0; i-- {
			donate := resp.Data.Donates[i]
			donationID := donate.ID // Use Reymit's donation ID as our primary key

			// Convert float timestamps and amounts to int64
			createdAt := int64(donate.Time)
			amount := int64(donate.Amount) // Assuming amount is in the smallest unit or int64 is sufficient

			slog.Debug("Processing donation", "id", donationID, "name", donate.Name, "amount", donate.Amount, "currency", donate.Currency, "time", donate.Time)

			// Check if this donation ID already exists in the database
			// We attempt to create it and handle the constraint error
			_, err := dbClient.ProcessedDonate.
				Create().
				SetID(donationID).
				SetDisplayName(donate.Name).
				SetCreatedAt(createdAt).
				SetAmount(amount).
				SetCurrency(donate.Currency).
				Save(ctx)

			if err != nil {
				if ent.IsConstraintError(err) {
					slog.Debug("Donation already processed", "id", donationID)
					// If this donation is already processed, assume older ones in this batch are too
					continue
				}
				slog.Error("Failed to save new donation to database", "id", donationID, "error", err)
				// Continue processing other donations in the batch even if one fails to save
				continue
			}

			slog.Info("New donation processed and saved", "id", donationID, "name", donate.Name, "amount", amount, "currency", donate.Currency)

			// Notify Telegram about the new donation
			// Fetch the saved entity to ensure we have the full struct if needed later
			savedDonation, err := dbClient.ProcessedDonate.Get(ctx, donationID)
			if err != nil {
				slog.Error("Failed to retrieve saved donation for notification", "id", donationID, "error", err)
				// Continue even if notification fails
				continue
			}

			err = bot.NotifyNewDonation(savedDonation)
			if err != nil {
				slog.Error("Failed to send Telegram notification", "id", donationID, "error", err)
				// Continue even if notification fails
			}
		}

		// Wait for 30 seconds before the next fetch
		time.Sleep(10 * time.Second)
	}
}
