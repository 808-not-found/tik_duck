package userservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func UserPublishActionService(
	ctx context.Context,
	req *user.PublishActionRequest,
) (*user.PublishActionResponse, error) {
	resp := user.PublishActionResponse{
		StatusCode: 0,
	}
	// 用户鉴权
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = 1024
		return &resp, err
	}
	myID := claims.ID
	// 写入数据
	err = db.UserPublishAction(ctx, myID, req.FilePath, req.CoverPath, req.Title)
	if err != nil {
		resp.StatusCode = 1025
		return &resp, err
	}
	// 成功返回
	return &resp, nil
}
