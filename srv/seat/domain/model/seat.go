package model

import "gorm.io/gorm"

type Seat struct {
	ID         uint64 `gorm:"primary_key;not_null;auto_increment"`
	Tag        uint64 `gorm:"default:0" json:"updated_at"`
	SeatNo     string `gorm:"not null" json:"seat_no"`
	CarNo      uint32 `gorm:"not null" json:"car_no"`
	SeatType   uint32 `gorm:"not null" json:"seat_type"`
	SeatStatus uint64 `gorm:"not null" json:"seat_status"`
	TrainID    uint64 `gorm:"not null" json:"train_id"`
}

func (s *Seat) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("Tag", s.Tag+1)
	return
}
