package handler

import (
	"context"
	"errors"
	"github.com/mamachengcheng/12306/srv/user/domain/model"
	"github.com/mamachengcheng/12306/srv/user/domain/service"
	"github.com/mamachengcheng/12306/srv/user/proto/user"
	"log"
	"regexp"
	"strconv"
	"time"
)

type User struct {
	UserDataService service.IUserDataService
}

func (u *User) Register(ctx context.Context, in *user.RegisterRequest, out *user.RegisterReply) error {
	out = &user.RegisterReply{
		IsSuccess: true,
		Msg:       "注册成功",
	}

	sex, birthday, err := ParseIdentityCard(in.Certificate)

	if err != nil {
		out.IsSuccess = false
		out.Msg = "身份证解析失败"
		return err
	}

	//passenger := model.Passenger{
	//	Name:        in.Name,
	//	Sex:         sex,
	//	Birthday:    birthday,
	//	Certificate: in.Certificate,
	//	MobilePhone: in.MobilePhone,
	//}

	user := &model.User{
		Username:    in.Username,
		Email:       in.Email,
		MobilePhone: in.MobilePhone,
		Password:    in.Password,
		Passengers: []model.Passenger{
			{
				Name:        in.Name,
				Sex:         sex,
				Birthday:    birthday,
				Certificate: in.Certificate,
				MobilePhone: in.MobilePhone,
			},
		},
	}

	//passengerID, err := u.UserDataService.AddPassenger(&passenger)

	//if err != nil {
	//	out.IsSuccess = false
	//	out.Msg = "注册失败"
	//	return err
	//}

	//user.UserInformationID = uint(passengerID)

	_, err = u.UserDataService.AddUser(user)

	if err != nil {
		out.IsSuccess = false
		out.Msg = "注册失败"
		return err
	}

	passengerID, _ := u.UserDataService.FindPassengerByCertificate(in.Certificate)
	//log.Printf("%v", passengerID)
	_, _ = u.UserDataService.UpdateUserPassengerID(user, uint(passengerID))

	return nil
}

func (u *User) Login(ctx context.Context, in *user.LoginRequest, out *user.LoginReply) error {
	out = &user.LoginReply{
		IsSuccess: true,
		Msg:       "登陆成功",
	}

	isSuccess, err := u.UserDataService.CheckPassword(in.Username, in.Password)
	log.Printf("handler----------")
	log.Printf("%v", isSuccess)
	log.Printf("%v", err)
	log.Printf("%v", out)

	if err != nil || !isSuccess {
		out.IsSuccess = isSuccess
		out.Msg = "登陆失败"
		log.Printf("%v", err)
		return err
	}

	return nil
}

func (u *User) QueryUserInformation(ctx context.Context, request *user.QueryUserInformationRequest, reply *user.QueryUserInformationReply) error {
	panic("implement me")
}

func (u *User) UpdatePassword(ctx context.Context, request *user.UpdatePasswordRequest, reply *user.UpdatePasswordReply) error {
	panic("implement me")
}

func (u *User) AddRegularPassenger(ctx context.Context, request *user.AddRegularPassengerRequest, reply *user.AddRegularPassengerReply) error {
	panic("implement me")
}

func (u *User) QueryRegularPassengers(ctx context.Context, request *user.QueryRegularPassengersRequest, reply *user.QueryRegularPassengersReply) error {
	panic("implement me")
}

func (u *User) UpdateRegularPassenger(ctx context.Context, request *user.UpdateRegularPassengerRequest, reply *user.UpdateRegularPassengerReply) error {
	panic("implement me")
}

func (u *User) DeleteRegularPassenger(ctx context.Context, request *user.DeleteRegularPassengerRequest, reply *user.DeleteRegularPassengerReply) error {
	panic("implement me")
}

func ParseIdentityCard(identityCard string) (sex bool, birthday time.Time, err error) {
	regular := "^\\d{6}(\\d{8})\\d{2}(\\d)[0-9X]$"
	reg := regexp.MustCompile(regular)

	result := reg.FindStringSubmatch(identityCard)

	if len(result) != 3 {
		return false, time.Time{}, errors.New("ParsingFailed")
	}

	sexNumber, err := strconv.Atoi(result[2])

	if err != nil {
		return false, time.Time{}, err
	}

	if sexNumber%2 == 1 {
		sex = true
	} else {
		sex = false
	}

	const format = "2006-01-02"
	birthday, err = time.Parse(format, result[1][:4]+"-"+result[1][4:6]+"-"+result[1][6:])

	return sex, birthday, err
}
