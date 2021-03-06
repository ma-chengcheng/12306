package service

import (
	"errors"
	"github.com/mamachengcheng/12306/srv/user/domain/model"
	"github.com/mamachengcheng/12306/srv/user/domain/respository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IUserDataService interface {
	AddUser(*model.User) (int64, error)
	UpdateUserPassengerID(user *model.User, id uint) (int64, error)
	CheckPassword(username string, pwd string) (isOk bool, err error)
	AddPassenger(passenger *model.Passenger) (int64, error)
	FindPassengerByCertificate(certificate string) (int64, error)
}

func NewUserDataService(userRepository respository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

type UserDataService struct {
	UserRepository respository.IUserRepository
}

func (u *UserDataService) UpdateUserPassengerID(user *model.User, id uint) (int64, error) {
	return u.UserRepository.UpdateUserPassengerID(user, id)
}

func (u *UserDataService) FindPassengerByCertificate(certificate string) (int64, error) {
	//panic("implement me")
	passenger, err := u.UserRepository.FindPassengerByCertificate(certificate)
	if err != nil {
		return -1, err
	}
	return passenger.ID, err
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

func (u *UserDataService) AddPassenger(passenger *model.Passenger) (int64, error) {
	//panic("implement me")
	return u.UserRepository.CreatePassenger(passenger)
}

func ValidatePassword(userPassword string, hashed string) (isSuccess bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		log.Printf("%v\n", err)
		return false, errors.New("IncorrectPassword")
	}
	return true, nil
}

func (u *UserDataService) CheckPassword(username string, password string) (isSuccess bool, err error) {
	user, err := u.UserRepository.FindUserByUsername(username)
	log.Printf("CheckPassword----------")
	log.Printf("%v", user)
	log.Printf("%v", err)

	if err != nil {
		log.Printf("%v\n", err)
		return false, err
	}

	return ValidatePassword(password, user.Password)
}
