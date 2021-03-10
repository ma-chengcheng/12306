package respository

import (
	"github.com/mamachengcheng/12306/srv/user/domain/model"
	"gorm.io/gorm"
	"strconv"
)

type IUserRepository interface {
	InitTable() error
	CreateUser(user *model.User) (int64, error)
	UpdateUserPassengerID(user *model.User, id uint) (int64, error)
	UpdateUserPassword(username string, pwdByte string) error
	FindUserByUsername(string) (*model.User, error)
	FindUserByEmail(string) (*model.User, error)
	FindUserByMobilePhone(string) (*model.User, error)
	CreatePassenger(passenger *model.Passenger, username string) (int64, error)
	FindPassengerByCertificate(certificate string) (*model.Passenger, error)
	FindPassengerByID(id uint) (*model.Passenger, error)
	FindPassengersByUsername(username string) (*[]model.Passenger, error)
	UpdatePassenger(passengerID string, mobilePhone string, passengerType string) error
	DeletePassenger(passengerID string) error
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDB: db}
}

type UserRepository struct {
	mysqlDB *gorm.DB
}

func (u *UserRepository) DeletePassenger(passengerID string) error {
	//panic("implement me")
	id, _ := strconv.ParseUint(passengerID, 10, 64)
	return u.mysqlDB.Delete(&model.Passenger{}, id).Error
}

func (u *UserRepository) UpdatePassenger(passengerID string, mobilePhone string, passengerType string) error {
	//panic("implement me")
	id, _ := strconv.ParseUint(passengerID, 10, 64)
	totype, _ := strconv.ParseUint(passengerType, 10, 64)
	passenger, err := u.FindPassengerByID(uint(id))
	if err != nil {
		return err
	}
	return u.mysqlDB.Model(passenger).Updates(model.Passenger{
		PassengerType: uint8(totype),
		MobilePhone:   mobilePhone,
	}).Error
}

func (u *UserRepository) FindPassengersByUsername(username string) (*[]model.Passenger, error) {
	//panic("implement me")
	user, _ := u.FindUserByUsername(username)
	passengers := &[]model.Passenger{}
	u.mysqlDB.Where("user_id = ?", user.ID).Find(passengers)
	return passengers, u.mysqlDB.Where("user_id = ?", user.ID).Find(passengers).Error
}

func (u *UserRepository) UpdateUserPassword(username string, pwdByte string) error {
	//panic("implement me")
	return u.mysqlDB.Model(&model.User{}).Where("username = ?", username).Update("password", pwdByte).Error
}

func (u *UserRepository) FindPassengerByID(id uint) (*model.Passenger, error) {
	passenger := &model.Passenger{}
	return passenger, u.mysqlDB.Where("id = ?", id).First(passenger).Error
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

func (u *UserRepository) CreatePassenger(passenger *model.Passenger, username string) (int64, error) {
	//panic("implement me")
	//log.Println(username)
	user, _ := u.FindUserByUsername(username)
	passenger.UserID = uint64(user.ID)
	//log.Println(passenger.UserID)
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
