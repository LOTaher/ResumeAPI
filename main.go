package main

import (
	"log"

	"github.com/LOTaher/resumeapi/database"
	"github.com/LOTaher/resumeapi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
	db, err := database.InitDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := fiber.New()

	app.Use(logger.New())

	// Setup routes from routes.go
	routes.SetupRoutes(app, db)

	app.Listen(":3000")
}
