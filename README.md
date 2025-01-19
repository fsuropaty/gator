# Gator-go
Simple CLI blog aggregator tool
> A guided project from [Boot.Dev](https://www.boot.dev/courses/build-blog-aggregator-golang) course

## Installation
Follow these steps to run this program on your machine

### Prerequisites
 - [Postgresql](https://www.postgresql.org/) 
 - [Go](https://go.dev/)
### Install
```bash
go install github.com/fsuropaty/gator-go
```
### Setting up
1. Manually create a config file in your home directory `~/.gatorconfig.json` with the following content
```json
{
    "db_url": "postgres://username:password@localhost:5432/database_name"
}
```
## Usage
### Basic Commands
```bash
gator <command> [arguments]
```
### User Management
 - `gator users` - Display the list of registered users
 - `gator register <name>` - Register a user
 - `gator login <name>` - Login to a given name
 - `gator reset` - Reset the users table

 ### Feed Management
 - `gator addfeed <RSS Name> <url>` - Add RSS feed source to current user
 - `gator feeds` - Display the list of the added RSS feed
 - `gator follow <url>` - Follow the added RSS feed
 - `gator unfollow <url>` - Unfollow a RSS feed
 - `gator following` - Display the list of following RSS feed
 - `gator agg <time>` - Aggregate the RSS feed to create post (example: `gator agg 5m` for 5 minutes)
 - `gator browse <limit>` - Display list of post(s) (example `gator browse 10` to show 10 post)





