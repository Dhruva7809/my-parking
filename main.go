package main

import (
	"github.com/0xPiyush/my-parking/db"
	"github.com/0xPiyush/my-parking/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)

	app.Listen(":8000")
}
