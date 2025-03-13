package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Londrin/rss-aggregator/internal/config"
	"github.com/Londrin/rss-aggregator/internal/database"
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
