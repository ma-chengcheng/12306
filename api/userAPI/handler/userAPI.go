package handler

import (
	user "github.com/mamachengcheng/12306/srv/user/proto"
)

type UserAPI struct{
	UserService user.UserService
}
