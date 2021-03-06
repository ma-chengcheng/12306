package model

import (
	"time"
)

type Seat struct {
	ID        uint64 `gorm:"primary_key;not_null;auto_increment"`
	UpdatedAt time.Time `json:"updated_at"`

	SeatNo     string `gorm:"not null" json:"seat_no"`
	CarNumber  uint32 `gorm:"not null" json:"car_number"`
	SeatType   uint32 `gorm:"not null" json:"seat_type"`
	SeatStatus uint64 `gorm:"not null" json:"seat_status"`
	TrainID    uint64 `gorm:"not null" json:"train_id"`
}
