package main

import (
	"chat/app/chat/graph"
	"chat/pkg/api"
	"log"
)

func main() {
	err := api.StartServer(".env", graph.Router())
	if err != nil {
		log.Fatalln(err)
	}
}
