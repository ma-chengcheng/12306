package respository

import (
	"github.com/mamachengcheng/12306/srv/seat/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	InitTable() error
	FindSeatsByScheduleID() error
	UpdateSeats() error
}

func FindSeatsByScheduleID(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDB: db}
}

type TicketRepository struct {
	mysqlDB *gorm.DB
}

func (u *UserRepository) InitTable() error {
	return u.mysqlDB.AutoMigrate(&model.User{}, &model.Passenger{})
}

func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID, u.mysqlDB.Create(user).Error
}

