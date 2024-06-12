package handlers

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/sanjaisamson/Ticket-Booking-App/src/database"
	"github.com/sanjaisamson/Ticket-Booking-App/src/models"
)

func addBooking(c *fiber.Ctx) error {
	// db := database.DB.Db
	user := new(models.Booking)
	// Store the body in the user and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    err,
		})
	}
	return nil
}
