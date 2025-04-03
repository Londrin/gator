package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Londrin/gator/internal/config"
	"github.com/Londrin/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	curCfg, err := config.Read()
	if err != nil {
		fmt.Printf("Unable to read config(second attempt): %v", err)
	}

	db, err := sql.Open("postgres", curCfg.DBUrl)
	if err != nil {
		fmt.Printf("Unable to open DB: %v", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	defer db.Close()

	curState := &state{
		db:  dbQueries,
		cfg: &curCfg,
	}

	curCmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	curCmds.register("login", handlerLogin)
	curCmds.register("register", handlerRegister)
	curCmds.register("reset", handlerReset)
	curCmds.register("users", handlerListUsers)
	curCmds.register("agg", middlewareLoggedIn(handlerAgg))
	curCmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	curCmds.register("feeds", handlerListFeeds)
	curCmds.register("follow", middlewareLoggedIn(handlerFollowFeed))
	curCmds.register("following", middlewareLoggedIn(handlerFollowingFeed))
	curCmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	curCmds.register("browse", middlewareLoggedIn(handlerBrowse))
	curCmds.register("help", handlerHelp)

	if len(os.Args) < 2 {
		fmt.Printf("Input Error - Not enough arguments: %v\n", os.Args)
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := command{
		Name: cmdName,
		Args: cmdArgs,
	}

	err = curCmds.run(curState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
