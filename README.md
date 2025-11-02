# Gator (Boot.dev Aggregator)

A command-line RSS feed aggregator written in Go.

## Requirements

Before you begin, make sure you have the following installed:

- [Go](https://go.dev/dl/) **version 1.21 or later**
- [PostgreSQL](https://www.postgresql.org/download/)

## Commands
- `register <username>` - Create a new user
- `login <username>` - Set the current active user
- `reset` - Delete all users from the database
- `users` - Show all users and indicate the current logged-in user
- `addfeed <url>` - Add new RSS feed
- `feeds` - View all available feeds
- `follow <url>` - Follow a feed
- `following` - Show feeds that the current user is following
- `unfollow <url>` - Unfollow a feed
- `agg <interval>` - Continuously fetch feeds (e.g. 1m)
- `browse [limit]` - View recent posts from followed feeds

## Example Usage

```bash
# Register and log in
gator register alice
gator login alice

# Add and follow feeds
gator addfeed https://techcrunch.com/feed/
gator follow https://techcrunch.com/feed/

# Start aggregator (fetch every minute)
gator agg 1m

# Browse posts from followed feeds
gator browse 5
```

## Setup
1. Run migrations using Goose
2. Generate SQLC code
3. Build and run:
   ```bash
   go run .


## Repository

Push your project to GitHub and submit the repo link:  
```
https://github.com/tintinsn/Gator
```
