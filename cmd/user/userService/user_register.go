package userservice

import (
	"context"
	"time"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	"github.com/808-not-found/tik_duck/pkg/salt"
)

// ToDo:
// 错误码统一定义
// 错误解析

func NewUserRegisterService(ctx context.Context) context.Context {
	return ctx
}

func UserRegisterService(ctx context.Context, req *user.UserRegisterRequest) (int32, string, int64, string, error) {
	var statusCode int32
	NewSalt := salt.GenerateRandomSalt(salt.DefaultsaltSize)
	NewPassWord := salt.HashPassword(req.Password, NewSalt)
	userinfo := db.User{
		CreateTime:    time.Now(),
		Name:          req.Username,
		Password:      NewPassWord,
		Salt:          NewSalt, // ToDo : 加盐
		FollowCount:   0,
		FollowerCount: 0,
	}
	err := db.CreateUser(ctx, &userinfo)
	Userid := userinfo.ID
	if err != nil {
		statusCode = 1001
		return statusCode, "", 0, "", err
	}
	Token, err := jwt.GenToken(req.Username)
	if err != nil {
		statusCode = 1002
		return statusCode, "", 0, "", err
	}
	// 成功返回
	statusMsg := "success"
	return statusCode, statusMsg, Userid, Token, nil
}
