package respository

import (
	"github.com/mamachengcheng/12306/srv/user/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	InitTable() error
	CreateUser(user *model.User) (int64, error)
	FindUserByUsername(string) (*model.User, error)
	FindUserByEmail(string) (*model.User, error)
	FindUserByMobilePhone(string) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDB: db}
}

type UserRepository struct {
	mysqlDB *gorm.DB
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
