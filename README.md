# Gator  [![Go Report Card](https://goreportcard.com/badge/github.com/Londrin/gator)](https://goreportcard.com/report/github.com/Londrin/gator)

Gator is a command line RSS Feed **Aggregator**

A one-stop shop for managing your rss feeds with an automatic feed aggregation to keep you up to date!

## Features include:
- Automatic rss feed updating
- Create RSS feeds for individual users
- Follow RSS feeds others added
- Automatic feed following upon adding a feed
- Database managed feeds

## Requirements
### Go Installation
Install [Go](https://go.dev/doc/install).

### Postgres Installation
Install [Postgres](https://www.postgresql.org/).

### Config File
Add `~/.gatorconfig.json` to your home directory and populate with db config settings: (remove comments prior to saving)
```go
{
  "db_url": "connection_string_goes_here",    // replace "connection_string_goes_here" with your postgres db string: "postgres://username:password@localhost:5432/gator"
  "current_user_name": "username_goes_here"   // replace "username_goes_here" with a "username": "fred"
}
```

## Install Gator
There are a few routes you can go but a couple include:

### Clone the repo and run go install in the project directory
```bash
git clone https://github.com/Londrin/gator.git
cd gator
go install
```

### After installing Go above, use go install
```bash
go install github.com/Londrin/gator@latest
```

### Download the release
You can find the latest download of [gator in releases](https://github.com/Londrin/gator/releases)

## Operating Gator
```bash
gator `command` [args...]
```

## Commands
- register `user`
  - Registers a new user
```bash
gator register fred
```

- login `user`
  - Logs in as a registered user
```bash
gator login george
```

- addfeed `title` `url`
  - Adds & follows a feed
```bash
gator addfeed hackernews https://hnrss.org/newest 
```

- users
  - Displays all users
```bash
gator users
```

- feeds
  - Displays all feeds
```bash
gator feeds
```

- follow `feedUrl`
  - Follows an existing feed
```bash
gator follow https://hnrss.org/newest
```

- unfollow `feedUrl`
  - Unfollows an existing feed
```bash
gator unfollow https://hnrss.org/newest
```

- following
  - Displays all feeds followed by current user
```bash
gator following
```

- agg `timestring`
  - Aggregates posts every given duration 
    - Valid time units are `"s" (seconds), "m" (minutes), "h" (hours)`
```bash
gator agg 1h
```

- browse `postsCount`
  - Displays the latest `postCount` posts saved from your feeds, in descending order.
```bash
gator browse 5
```

- reset
  - Clears all data (user, feed, follow and posts)
```bash
gator reset
```
