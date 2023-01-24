package userservice

import (
	"context"
	"time"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

// ToDo:
// 错误码统一定义
// 错误解析

func NewUserRegisterService(ctx context.Context) context.Context {
	return ctx
}

func UserRegisterService(ctx context.Context, req *user.UserRegisterRequest) (int32, string, int, string, error) {
	var statusCode int32
	userinfo := db.User{
		CreateTime:    time.Now(),
		Name:          req.Username,
		Password:      req.Password,
		Salt:          "", // ToDo : 加盐
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
