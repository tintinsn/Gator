# Gator (Boot.dev Aggregator)

A command-line RSS feed aggregator written in Go.

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

## Setup
1. Run migrations using Goose
2. Generate SQLC code
3. Build and run:
   ```bash
   go run .
