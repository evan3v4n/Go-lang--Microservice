package main

import (
	"log"
	"Go_lang_Microservice/api/routes"
	"Go_lang_Microservice/config"
	"Go_lang_Microservice/db"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	
	if err := db.InitPostgres(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	if err := db.InitRedis(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	r := routes.SetupRouter()

	port := config.Get("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
