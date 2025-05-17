# DonateNotifier

A Go application that monitors Reymit donations, stores them in a PostgreSQL database using Ent, and sends notifications to a Telegram group.

## Features

- Fetches latest donations from the Reymit API.
- Stores processed donations in a database to avoid duplicate notifications.
- Sends formatted notifications to a specified Telegram group.
- Configurable via environment variables.

## Prerequisites

- Go (version 1.20 or higher recommended)
- PostgreSQL database
- Reymit API Token
- Telegram Bot Token
- Telegram Group ID

## Setup

1.  **Clone the repository:**
    ```bash
    git clone <repository_url>
    cd DonateNotifier
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Configure environment variables:**
    Create a `.env` file in the project root based on the `.env.sample`.
    ```bash
    cp .env.sample .env
    ```
    Edit the `.env` file and fill in your `REYMIT_TOKEN`, `BOT_TOKEN`, `GROUP_ID`, and `POSTGRESQL_DSN`.
## Running the Application

```bash
go run main.go
```

The application will start fetching donations periodically (every 30 seconds by default) and send notifications for new donations to your configured Telegram group.

## Technologies Used

- Go
- Ent (ORM for database interaction)
- PostgreSQL
- Reymit API
- Telebot (Telegram Bot API library)
```
