package main

import (
	"backend/api"
	"backend/ent"
	"backend/utils"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("load serverConfig failed: ", err)
	}

	client, err := ent.Open("postgres", config.Dsn)
	if err != nil {
		log.Fatal("connect to db failed: ", err)
	}
	defer client.Close()

	server, err := api.NewServer(config, client)
	if err != nil {
		log.Fatal("new server intance failed: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("start server failed: ", err)
	}
}
