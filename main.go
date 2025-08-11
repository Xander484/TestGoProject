package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"myapp/db"
	"myapp/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("cannot.read.env.file")
	}
}

func main() {
	// Connect to DB
	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbConn.Close()

	// Get port from env or default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Create Gin router with DB
	r := routes.NewRouter(dbConn)

	fmt.Println("Server is running on port " + port + "...")

	// Run Gin server
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server error: ", err)
	}
}
