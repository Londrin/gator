package main

import "fmt"

func handlerHelp(s *state, cmd command) error {
	fmt.Println("Gator Help!")
	fmt.Println()

	fmt.Println("Usage: gator <command> [args..]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println()
	fmt.Println("register <user_name>	|	registers a new user")
	fmt.Println()
	fmt.Println("login <user_name>	|	logs in as a registered user")
	fmt.Println()
	fmt.Println("addfeed <name> <url>	|	adds and follows a new rss feed 'addfeed techcrunch https://techcrunch.com/feed/'")
	fmt.Println()
	fmt.Println("users			|	displays all registered users")
	fmt.Println()
	fmt.Println("feeds			|	displays all feeds in the database")
	fmt.Println()
	fmt.Println("follow <url>		|	follows an existing feed in the database")
	fmt.Println()
	fmt.Println("unfollow <url>		|	unfollows an existing feed in the database")
	fmt.Println()
	fmt.Println("agg <timestring>	|	automatically aggregates posts from feeds in the database on given timestring `e.g. 10s, 5m, 3h`")
	fmt.Println()
	fmt.Println("browse <post_count>	|	displays the latest `post_count` posts saved from your feeds, in newest first order (DESCENDING)")
	fmt.Println()
	fmt.Println("reset			|	completely wipes the database of all users, feeds, follows and posts")

	return nil
}
