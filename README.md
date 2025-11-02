# Gator (Boot.dev Aggregator)

A command-line RSS feed aggregator written in Go.

## Commands
- `addfeed <url>` - Add new RSS feed
- `follow <url>` - Follow a feed
- `agg <interval>` - Continuously fetch feeds (e.g. 1m)
- `browse [limit]` - View recent posts

## Setup
1. Run migrations using Goose
2. Generate SQLC code
3. Build and run:
   ```bash
   go run .
