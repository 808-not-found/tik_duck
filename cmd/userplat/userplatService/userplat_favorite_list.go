package userplatservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/userplat/dal/db"
	"github.com/808-not-found/tik_duck/cmd/userplat/pack"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func UserFavoriteListService(
	ctx context.Context,
	req *userplat.FavoriteListRequest,
) (*userplat.FavoriteListResponse, error) {
	var resp userplat.FavoriteListResponse

	// 用户鉴权
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = 1007
		return &resp, nil
	}

	// 检查登录状态
	myID := claims.ID
	if myID == 0 {
		resp.StatusCode = 1008
		return &resp, err
	}
	// 查询数据库
	var dbVideos []*db.Video
	dbVideos, err = db.GetFavoriteList(ctx, req.UserId)
	if err != nil {
		resp.StatusCode = 1006
		return &resp, err
	}
	// 数据封装
	rpcVideos, err := pack.Videos(ctx, dbVideos, myID)
	if err != nil {
		resp.StatusCode = 1007
		return &resp, err
	}
	resp.VideoList = rpcVideos

	return &resp, nil
}
