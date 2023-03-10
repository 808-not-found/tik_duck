package userservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/cmd/user/pack"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func UserPublishListService(ctx context.Context, req *user.PublishListRequest) (*user.PublishListResponse, error) {
	// 创建返回参数
	resp := user.PublishListResponse{
		StatusCode: 0,
	}

	// 用户鉴权
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = 1026
		return &resp, nil
	}
	myID := claims.ID
	// 获取数据
	dbVideos, err := db.UserPublishList(ctx, int32(req.UserId))
	if err != nil {
		resp.StatusCode = 1027
		return &resp, nil
	}
	// 封装数据
	rpcVideos, err := pack.Videos(ctx, dbVideos, myID)
	if err != nil {
		resp.StatusCode = 1028
		return &resp, nil
	}
	resp.VideoList = rpcVideos

	// 成功返回
	return &resp, nil
}
