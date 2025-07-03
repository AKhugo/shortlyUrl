package main

import (
	"log"
	"shortlyurl/internal/config"
	"shortlyurl/internal/database"

	"github.com/joho/godotenv"
)

func main() {

	// load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Charger la configuration après le .env
	config.LoadConfig()

	// init database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// get underlying sql.DB pour fermer la connexion
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get underlying sql.DB:", err)
	}
	defer sqlDB.Close()

	log.Println("Database connected successfully!")
}
