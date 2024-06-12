package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type Booking struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;"`
	Username  string    `json:"username"`
	UserId    string    `gorm:"type:uuid;"`
	EventName string    `json:"eventname"`
	Seats     string    `json:"seats"`
}

// Users struct
type Bookings struct {
	Bookings []Booking `json:"tickets"`
}

func (booking *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	booking.ID = uuid.New()
	return
}
