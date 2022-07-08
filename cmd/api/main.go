package main

import (
	"log"
	"meli_test/internal/api/routes"
	"os"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// @title MeLi Test
// @version 1.0
// @description This is the main to microservice.
// @contact.name Gerzon Bautista Gonzalez
// @contact.email gerzonbautista@gmail.com
// @BasePath /
func main() {
	var port string
	app := fiber.New(fiber.Config{
		ReadBufferSize: 9000,
	})

	routes.SetupRoutes(app)
	app.Get("/swagger/*", swagger.Handler)

	if port = os.Getenv("HTTP_PORT"); port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
