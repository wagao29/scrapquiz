package main

import (
	"context"
	"log"

	"scrapquiz/config"
	"scrapquiz/infrastructure/mysql/db"
	"scrapquiz/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.GetConfig()
	log.Printf("config: %+v", conf)

	db.NewMainDB(conf.DB)

	server.Run(ctx, conf)
}
