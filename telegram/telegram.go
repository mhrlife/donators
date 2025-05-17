package telegram

import (
	"DonateNotifier/ent"
	"DonateNotifier/utils"
	"fmt"
	"gopkg.in/telebot.v4"
	"log/slog"
)

// Bot represents the Telegram bot client and configuration.
type Bot struct {
	client  *telebot.Bot
	groupID telebot.ChatID
}

// NewBot initializes and returns a new Telegram Bot instance.
func NewBot(token string, groupID int64) (*Bot, error) {
	pref := telebot.Settings{
		Token: token,
		// Add other settings if needed, e.g., Poller
		// Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("error creating bot: %w", err)
	}

	slog.Info("Telegram bot initialized", "username", b.Me.Username)

	return &Bot{
		client:  b,
		groupID: telebot.ChatID(groupID),
	}, nil
}

// SendMessage sends a message to the configured group ID.
func (b *Bot) SendMessage(text string, options ...any) error {
	_, err := b.client.Send(b.groupID, text, options...)
	if err != nil {
		return fmt.Errorf("error sending message to group %d: %w", b.groupID, err)
	}
	return nil
}

// NotifyNewDonation sends a formatted message to the Telegram group about a new donation.
func (b *Bot) NotifyNewDonation(donation *ent.ProcessedDonate) error {
	// Format amount with commas
	formattedAmount := utils.FormatAmountWithCommas(donation.Amount)

	// Translate currency to Farsi
	translatedCurrency := utils.TranslateCurrency(donation.Currency)

	// Escape display name and formatted amount for MarkdownV2
	escapedDisplayName := utils.EscapeMarkdown(donation.DisplayName)
	escapedAmount := utils.EscapeMarkdown(formattedAmount)
	escapedCurrency := utils.EscapeMarkdown(translatedCurrency)

	// Construct the Farsi message using MarkdownV2
	// Example: تشکر از [display_name] برای حمایت با ☕️ *[amount] [currency]*
	message := fmt.Sprintf(`☕️✨ ممنون از *%s* با *%s %s* منو به کلی قهوه‌ی تازه مهمون کرد مرسی رفیق، حمایتت واقعاً انرژی‌بخشه ❤️

[لیست کامل حامی‌ها](https://reymit.ir/mhrcode/donators)`,
		escapedDisplayName,
		escapedAmount,
		escapedCurrency,
	)

	slog.Info("Sending Telegram notification", "donation_id", donation.ID, "message", message)

	// Send the message using MarkdownV2 mode
	_, err := b.client.Send(b.groupID, message, &telebot.SendOptions{
		DisableWebPagePreview: true,
		ParseMode:             telebot.ModeMarkdownV2,
	})
	if err != nil {
		return fmt.Errorf("failed to send donation notification for ID %s: %w", donation.ID, err)
	}

	slog.Info("Telegram notification sent successfully", "donation_id", donation.ID)
	return nil
}

// StartPolling starts the Telegram bot polling for updates (if needed for other features).
// Currently, this bot only sends messages, so polling might not be necessary unless
// you plan to add command handlers or other interactive features.
func (b *Bot) StartPolling() {
	slog.Info("Starting Telegram bot polling")
	b.client.Start()
}

// StopPolling stops the Telegram bot polling.
func (b *Bot) StopPolling() {
	slog.Info("Stopping Telegram bot polling")
	b.client.Stop()
}
