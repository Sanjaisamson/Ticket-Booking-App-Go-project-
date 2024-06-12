// handlers_test.go
package handler_test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/sanjaisamson/Ticket-Booking-App/mock"

	// "github.com/sanjaisamson/Ticket-Booking-App/src/handlers"
	"github.com/sanjaisamson/Ticket-Booking-App/src/models"
	// "github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	mockCtx := new(mock.MockCtx)
	mockFiberCtx := new(mock.MockFiberCtx)
	mockDB := new(mock.MockDB)

	newUser := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "testpassword",
	}

	mockCtx.On("BodyParser", newUser).Return(nil)
	mockDB.On("Where", "email = ?", newUser.Email).Return(&gorm.DB{})
	mockDB.On("First", &models.User{}).Return(&gorm.DB{})
	mockDB.On("Create", newUser).Return(&gorm.DB{})
	mockFiberCtx.On("Status", fiber.StatusCreated).Return(&fiber.Ctx{})
	mockCtx.On("JSON", fiber.Map{"status": "success", "message": "User has created", "data": newUser}).Return(nil)

	mockCtx.FiberCtx = mockFiberCtx // Assign the MockFiberCtx instance to the FiberCtx field

	// err := handlers.CreateUser(mockCtx)
	// assert.NoError(t, err)

	mockCtx.AssertExpectations(t)
	mockDB.AssertExpectations(t)
	mockFiberCtx.AssertExpectations(t)
}
