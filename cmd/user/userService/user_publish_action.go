package userservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func UserPublishActionService(ctx context.Context, req *user.PublishActionRequest) (int32, string, error) {
	var statusCode int32
	var statusMsg string

	// 取出请求信息
	token := req.Token
	filePath := req.FilePath
	coverPath := req.CoverPath
	title := req.Title

	// 用户鉴权
	claims, err := jwt.ParseToken(token)
	if err != nil {
		statusCode = 1103
		return statusCode, statusMsg, err
	}
	username := claims.Username
	userInfo, err := db.QueryUser(ctx, username)
	if err != nil {
		statusCode = 1104
		return statusCode, statusMsg, err
	}
	userID := userInfo.ID

	// 向数据库中写入数据
	err = db.UserPublishAction(ctx, userID, filePath, coverPath, title)
	if err != nil {
		statusCode = 1105
		return statusCode, statusMsg, err
	}

	// 成功返回
	statusMsg = consts.Success
	return statusCode, statusMsg, nil
}
