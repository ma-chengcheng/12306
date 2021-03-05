package model

type Seat struct {
	ID         int64  `gorm:"primary_key;not_null;auto_increment"`

	SeatNo     string `gorm:"not null" json:"seat_no"`
	CarNumber  uint   `gorm:"not null" json:"car_number"`
	SeatType   uint   `gorm:"not null" json:"seat_type"`
	SeatStatus uint64 `gorm:"not null" json:"seat_status"`
	TrainID    uint   `gorm:"not null" json:"train_id"`
}
