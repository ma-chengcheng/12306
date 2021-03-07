package respository

import (
	"github.com/mamachengcheng/12306/srv/user/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	InitTable() error
	CreateUser(user *model.User) (int64, error)
	UpdateUserPassengerID(user *model.User, id uint) (int64, error)
	FindUserByUsername(string) (*model.User, error)
	FindUserByEmail(string) (*model.User, error)
	FindUserByMobilePhone(string) (*model.User, error)
	CreatePassenger(passenger *model.Passenger) (int64, error)
	FindPassengerByCertificate(certificate string) (*model.Passenger, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDB: db}
}

type UserRepository struct {
	mysqlDB *gorm.DB
}

func (u *UserRepository) UpdateUserPassengerID(user *model.User, id uint) (int64, error) {
	//panic("implement me")
	//utils.MysqlDB.Model(&models.User{}).Where("username= ?", claims.Username).Update("password", data.Password)
	return user.ID, u.mysqlDB.Model(&model.User{}).Where("id= ?", user.ID).Update("user_information_id", id).Error
}

func (u *UserRepository) FindPassengerByCertificate(certificate string) (*model.Passenger, error) {
	//panic("implement me")
	passenger := &model.Passenger{}
	return passenger, u.mysqlDB.Where("certificate = ?", certificate).First(passenger).Error
}

func (u *UserRepository) CreatePassenger(passenger *model.Passenger) (int64, error) {
	//panic("implement me")
	return passenger.ID, u.mysqlDB.Create(passenger).Error
}

func (u *UserRepository) InitTable() error {
	return u.mysqlDB.AutoMigrate(&model.User{}, &model.Passenger{})
}

func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID, u.mysqlDB.Create(user).Error
}

func (u *UserRepository) FindUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	return user, u.mysqlDB.Where("username = ?", username).First(user).Error
}

func (u *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	return user, u.mysqlDB.Where("email = ?", email).First(user).Error
}

func (u *UserRepository) FindUserByMobilePhone(mobilePhone string) (*model.User, error) {
	user := &model.User{}
	return user, u.mysqlDB.Where("mobilePhone = ?", mobilePhone).First(user).Error
}
