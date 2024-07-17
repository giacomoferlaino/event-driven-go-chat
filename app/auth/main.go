package main

import (
	"chat/app/auth/graph"
	"chat/pkg/env"
	"log"
)

func main() {
	err := env.Init(".env")
	if err != nil {
		log.Fatalln("Error initializing environment:", err)
	}

	router := graph.Router()

	port := env.Port()
	err = router.Run(":" + port)
	log.Fatalf("Error starting HTTP server:\n%s", err)
}
