package userservice

import (
	"context"
	"errors"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	"github.com/808-not-found/tik_duck/pkg/salt"
	"gorm.io/gorm"
)

func UserLoginService(ctx context.Context, req *user.UserLoginRequest) (int32, string, int64, string, error) {
	var statusCode int32
	// 取出请求中的两个信息
	reqUsername := req.Username
	reqPassword := req.Password

	// 1.对比数据库内容
	// 比对失败
	userinfo, err := db.QueryUser(ctx, reqUsername)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		statusCode = 1005
		return statusCode, "", 0, "", err
	case err != nil:
		statusCode = 1004
		return statusCode, "", 0, "", err
	}
	if !salt.PasswordsMatch(userinfo.Password, reqPassword, userinfo.Salt) {
		statusCode = 1006
		return statusCode, "", 0, "", err
	}

	// 2. 创建token
	Token, err := jwt.GenToken(req.Username)
	if err != nil {
		statusCode = 1007
		return statusCode, "", 0, "", err
	}

	statusMsg := consts.Success
	Userid := userinfo.ID

	return statusCode, statusMsg, Userid, Token, nil
}
