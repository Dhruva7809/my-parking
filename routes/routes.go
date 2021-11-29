package routes

import (
	"github.com/0xPiyush/my-parking/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user", controllers.User)
	app.Get("/api/locations", controllers.Locations)
	app.Get("/api/bookings", controllers.Bookings)
	app.Post("/api/book", controllers.AddBooking)
}
