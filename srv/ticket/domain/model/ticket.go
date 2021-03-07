package model

type Ticket struct {
	ID          uint64 `gorm:"primary_key;not_null;auto_increment"`
	SeatID      uint64 `gorm:"not_null" json:"seat_id"`
	PassengerID uint64 `gorm:"not_null" json:"passenger_id"`
	OrderID     uint64 `gorm:"not_null" json:"order_id"`
	Price       uint64 `gorm:"not_null" json:"price"`
}
