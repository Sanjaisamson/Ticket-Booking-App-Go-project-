package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
	"github.com/sanjaisamson/Ticket-Booking-App/src/database"
	"github.com/sanjaisamson/Ticket-Booking-App/src/routers"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	routers.SetupRoutes(app)
	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendString("this route is not available") // => 404 "Not Found"
	})
	app.Listen(":3000")
}
