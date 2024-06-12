package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sanjaisamson/Ticket-Booking-App/src/handlers"
	"github.com/sanjaisamson/Ticket-Booking-App/src/middleware"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/user")
	// routes
	v1.Post("/create", handlers.CreateUser)
	v1.Post("/login", handlers.LoginUser)
	v1.Post("/logout", middleware.AccessTokenVerification, handlers.Logout)
	v1.Get("/refresh", middleware.RefreshToken)
}
