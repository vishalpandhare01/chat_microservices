package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/vishalpandhare01/myschool_chat_microservices/initializer"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load env ")
	}
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000/, http://localhost:3001/, http://localhost:3002/",
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
	}))
	initializer.ConnectDB()
	initializer.MigrateTables()
	internal.SetupRoutes(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}
