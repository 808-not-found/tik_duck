package userservice

import (
	"context"
	"time"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	"github.com/808-not-found/tik_duck/pkg/salt"
)

func UserRegisterService(ctx context.Context, req *user.UserRegisterRequest) (int32, string, int64, string, error) {
	var statusCode int32
	NewSalt := salt.GenerateRandomSalt(salt.DefaultsaltSize)
	NewPassWord := salt.HashPassword(req.Password, NewSalt)
	userinfo := db.User{
		CreateTime:    time.Now(),
		Name:          req.Username,
		Password:      NewPassWord,
		Salt:          NewSalt,
		FollowCount:   0,
		FollowerCount: 0,
	}
	Exist, _ := db.CheckUserExist(ctx, req.Username)
	if Exist {
		statusCode = 1012
		return statusCode, "", 0, "", nil
	}
	err := db.CreateUser(ctx, &userinfo)
	if err != nil {
		statusCode = 1002
		return statusCode, "", 0, "", err
	}
	Userid, _ := db.GetUserID(ctx, req.Username)
	// 加了用户ID信息
	Token, err := jwt.GenToken(req.Username, Userid)
	if err != nil {
		statusCode = 1003
		return statusCode, "", 0, "", err
	}
	// 成功返回
	statusMsg := consts.Success
	return statusCode, statusMsg, Userid, Token, nil
}
