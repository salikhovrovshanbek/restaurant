package main

import (
	"github.com/gokurs/Projects/restaurant/config"
	"github.com/gokurs/Projects/restaurant/connect"
	"github.com/gokurs/Projects/restaurant/repository/postgres"
	"github.com/gokurs/Projects/restaurant/server"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln("failed to load:", err)
	}
	db, err := connect.Connect(cfg)
	if err != nil {
		log.Fatalln("failed to connect:", err)
	}
	repo := postgres.New(db)
	server.NewRoutor(repo, cfg)
}
