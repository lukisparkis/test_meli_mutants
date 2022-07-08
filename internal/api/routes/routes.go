package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"meli_test/internal/api/controllers/mutants"
	"meli_test/internal/api/controllers/stats"
)

// SetupRoutes is a route management of the app
func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome, This is a Test for MeLi")
	})

	m := app.Group("/mutant")
	{
		m.Post("/", mutants.Check)
	}

	s := app.Group("/stats")
	{
		s.Get("/", stats.GetAllChecksPrevious)
	}
}
