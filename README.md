# gator

## Overview
gator is a blog aggregator CLI service that fetches posts from a provided RSS feed on the internet. You can register, login, add feeds, follow/unfollow feeds, view feeds, view following feeds, and more all in a CLI.


## Prerequistes
You must have Postgres and Go's toolchain to run this program. You can download them from their official websites or through command line based on your OS and preference.


## Installation
Install the CLI tool using
```bash
go install github.com/haidar1337/gator@latest
```

After installing the tool, create a config file `.gatorconfig.json` in your home directory with the following structure:
```json
{
    "db_url": "postgres_connection_string",
    "current_user_name": "anything"
}
```

## Usage
To start you can register to the CLI using:

```bash
gator register <username>
```

Supported Commands:
- gator register \<username>
- gator login \<username>
- gator addfeed <feed_name> <feed_url>
- gator following
- gator feeds
- gator follow <feed_url>
- gator unfollow <feed_url>
- gator agg [time_between_requests]
- gator browse [limit]
