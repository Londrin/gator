package main

import (
	"fmt"
	"os"

	"github.com/Londrin/rss-aggregator/internal/config"
)

func main() {
	curCfg, err := config.Read()
	if err != nil {
		fmt.Printf("Unable to read config(second attempt): %v", err)
	}
	fmt.Printf("Reading config: %+v\n", curCfg)

	curState := state{
		cfg: &curCfg,
	}
	curCmds := commands{
		cmds: make(map[string]func(*state, command) error),
	}
	curCmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Printf("Input Error - Not enough arguments: %v\n", os.Args)
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	err = curCmds.run(&curState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	curCfg, err = config.Read()
	if err != nil {
		fmt.Printf("Unable to read config(second attempt): %v", err)
	}

	fmt.Printf("Reading config again: %+v\n", curCfg)
}
