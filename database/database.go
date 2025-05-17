package database

import (
	"context"
	"fmt"
	"os"

	"DonateNotifier/ent"

	_ "github.com/lib/pq"
)

func InitDB() (*ent.Client, error) {
	dsn := os.Getenv("POSTGRESQL_DSN")
	if dsn == "" {
		return nil, fmt.Errorf("POSTGRESQL_DSN environment variable not set")
	}

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %w", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		client.Close()
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}

	return client, nil
}
