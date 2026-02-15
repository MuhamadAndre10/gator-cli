# Blog Aggregator

A CLI application to aggregate RSS feeds from various blogs. Built with Go and PostgreSQL.

## Features

- **User Management**: Register and login users
- **RSS Feed Aggregation**: Fetch and parse RSS feeds from any URL
- **Feed Management**: Add and manage RSS feed subscriptions

## Prerequisites

- Go 1.21+
- PostgreSQL
- [Goose](https://github.com/pressly/goose) - Database migration tool
- [sqlc](https://sqlc.dev/) - SQL code generator

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/blog-aggregator.git
cd blog-aggregator

# Install dependencies
make deps

# Run database migrations
make migrate-up

# Build the application
make build
```

## Configuration

Create a config file at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://user:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Or set the database URL via environment variable.

## Usage

### User Commands

```bash
# Register a new user
./bin/gator register <username>

# Login as a user
./bin/gator login <username>

# List all users
./bin/gator users

# Reset (delete all users)
./bin/gator reset
```

### Feed Commands

```bash
# Add a new RSS feed
./bin/gator addfeed "<feed_name>" "<feed_url>"

# Example
./bin/gator addfeed "Hacker News" "https://hnrss.org/newest"
```

### Aggregator Commands

```bash
# Fetch and display RSS feed content
./bin/gator agg
```

## Development

```bash
# Run migrations
make migrate-up

# Generate sqlc code after modifying SQL queries
make sqlc

# Run tests
make test

# Build for development
make dev
```

## Project Structure

```
blog-aggregator/
├── internal/
│   ├── command/     # CLI command handlers
│   ├── config/      # Configuration management
│   ├── database/    # Generated database code (sqlc)
│   └── rss/         # RSS feed fetching and parsing
├── sql/
│   ├── queries/     # SQL queries for sqlc
│   └── schema/      # Database migrations (goose)
├── main.go
├── Makefile
└── README.md
```

## License

MIT
