package main

import (
	"fmt"
	"log"

	"github.com/Londrin/rss-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Ready config: %+v\n", cfg)

	err = cfg.SetUser("jeff")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Ready config again: %+v\n", cfg)
}
