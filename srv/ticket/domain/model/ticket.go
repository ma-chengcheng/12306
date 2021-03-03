package model

type Ticket struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment"`

	SeatID      int64 `gorm:"not_null"`
	PassengerID int64 `gorm:"not_null"`
	OrderID     int64 `gorm:"not_null"`
}
