package model

import (
	"time"
)

type User struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment"`

	Username    string `gorm:"not null;unique" json:"username"`
	Email       string `gorm:"not null;unique" json:"email"`
	MobilePhone string `gorm:"not null;unique" json:"mobile_phone"`

	Password string `gorm:"not null" json:"password"`

	UserInformationID uint `json:"user_information_id"`

	Passengers []Passenger
}

type Passenger struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment"`

	Name                string    `gorm:"not null" json:"name"`
	CertificateType     uint8     `gorm:"default:0"  json:"certificate_type"`
	Sex                 bool      `gorm:"not null" json:"sex"`
	Birthday            time.Time `gorm:"not null" json:"birthday"`
	Country             string    `gorm:"default:中国CHINA" json:"country"`
	CertificateDeadline time.Time `gorm:"default:'9999-12-31 23:59:59'" json:"certificate_deadline"`
	Certificate         string    `gorm:"not null" json:"certificate"`
	PassengerType       uint8     `gorm:"default:0" json:"passenger_type"`
	MobilePhone         string    `gorm:"not null" json:"mobile_phone"`
	CheckStatus         uint8     `gorm:"default:0" json:"check_status"`
	UserStatus          uint8     `gorm:"default:0" json:"user_status"`

	UserID uint64
}
