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

	router := httpserver.GetRouter()
	err = router.Run()
	log.Fatalf("Error starting HTTP server:\n%s", err)
}
