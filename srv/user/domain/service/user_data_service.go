package service

import (
	"errors"
	"github.com/mamachengcheng/12306/srv/user/domain/model"
	"github.com/mamachengcheng/12306/srv/user/domain/respository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(*model.User) (int64, error)
	CheckPassword(username string, pwd string) (isOk bool, err error)
}

func NewUserDataService(userRepository respository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

type UserDataService struct {
	UserRepository respository.IUserRepository
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func (u *UserDataService) AddUser(user *model.User) (int64, error) {
	pwdByte, err := GeneratePassword(user.Password)

	if err != nil {
		return user.ID, err
	}

	user.Password = string(pwdByte)

	return u.UserRepository.CreateUser(user)
}

func ValidatePassword(userPassword string, hashed string) (isSuccess bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("IncorrectPassword")
	}
	return true, nil
}

func (u *UserDataService) CheckPassword(username string, password string) (isSuccess bool, err error) {
	user, err := u.UserRepository.FindUserByUsername(username)

	if err != nil {
		return false, err
	}

	return ValidatePassword(password, user.Password)
}
