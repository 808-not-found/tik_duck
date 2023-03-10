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
	var myID int64
	if req.Token == "" {
		myID = 0
	} else {
		claims, err := jwt.ParseToken(req.Token)
		if err != nil {
			resp.StatusCode = 2037
			return &resp, nil
		}
		myID = claims.ID
	}
	// 检查登录状态
	// myID := claims.ID
	// if myID == 0 {
	// 	resp.StatusCode = 1008
	// 	return &resp, err
	// }
	// 查询数据库
	dbVideos, err := db.GetFavoriteList(ctx, req.UserId)
	if err != nil {
		resp.StatusCode = 2038
		return &resp, err
	}
	// 数据封装
	rpcVideos, err := pack.Videos(ctx, dbVideos, myID)
	if err != nil {
		resp.StatusCode = 2039
		return &resp, err
	}
	resp.VideoList = rpcVideos

	return &resp, nil
}
