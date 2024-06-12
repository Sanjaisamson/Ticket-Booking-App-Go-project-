package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type Ticket struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;"`
	Username  string    `json:"username"`
	UserId    string    `gorm:"type:uuid;"`
	EventName string    `json:"eventname"`
	BookingID uuid.UUID `gorm:"type:uuid;"`
	SeatNo    string    `json:"seatNo"`
}

// Users struct
type Tickets struct {
	Tickets []Ticket `json:"tickets"`
}

func (ticket *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	ticket.ID = uuid.New()
	return
}
