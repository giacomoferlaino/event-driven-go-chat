package main

import (
	"chat/internal/app/auth/env"
	"chat/internal/app/auth/httpserver"
	"log"
)

func main() {
	err := env.Init()
	if err != nil {
		log.Fatalln("Error initializing environment:", err)
	}

	router := httpserver.Router()

	port := env.Port()
	err = router.Run(":" + port)
	log.Fatalf("Error starting HTTP server:\n%s", err)
}
