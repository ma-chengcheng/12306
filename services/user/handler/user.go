package handler

import (
	"context"
	"errors"
	"github.com/mamachengcheng/12306/services/user/domain/model"
	"github.com/mamachengcheng/12306/services/user/domain/service"
	"regexp"
	"strconv"
	"time"

	user "github.com/mamachengcheng/12306/services/user/proto"
)

type User struct {
	UserDataService service.IUserDataService
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

	_, err = u.UserDataService.AddUser(user)

	if err != nil {
		out.IsSuccess = false
		out.Msg = "注册失败"
		return err
	}

	return nil
}

func (u *User) Login(ctx context.Context, in *user.LoginRequest, out *user.LoginReply) error {
	out = &user.LoginReply{
		IsSuccess: true,
		Msg:       "登陆成功",
	}

	isSuccess, err := u.UserDataService.CheckPassword(in.Username, in.Password)

	if err != nil || !isSuccess {
		out.IsSuccess = isSuccess
		out.Msg = "登陆失败"
		return err
	}

	return nil
}

func (u *User) QueryUserInformation(context.Context, *user.QueryUserInformationRequest, *user.QueryUserInformationReply) error {
	return nil
}
func (u *User) UpdatePassword(context.Context, *user.UpdatePasswordRequest, *user.UpdatePasswordReply) error {
	return nil
}
func (u *User) AddRegularPassenger(context.Context, *user.AddRegularPassengerRequest, *user.AddRegularPassengerReply) error {
	return nil
}
func (u *User) QueryRegularPassengers(context.Context, *user.QueryRegularPassengersRequest, *user.QueryRegularPassengersReply) error {
	return nil
}
func (u *User) UpdateRegularPassenger(context.Context, *user.UpdateRegularPassengerRequest, *user.UpdateRegularPassengerReply) error {
	return nil
}
func (u *User) DeleteRegularPassenger(context.Context, *user.DeleteRegularPassengerRequest, *user.DeleteRegularPassengerReply) error {
	return nil
}
