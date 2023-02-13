package userplatservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/userplat/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func UserFavoriteActionService(
	ctx context.Context,
	req *userplat.FavoriteActionRequest,
) (*userplat.FavoriteActionResponse, error) {
	var resp userplat.FavoriteActionResponse

	// 用户鉴权
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = 1007
		return &resp, nil
	}
	// 获取必要信息
	// 1.获取登录用户ID
	// 2.获取当前视频ID
	// 检查是否点赞
	myID := claims.ID
	vdID := req.VideoId
	actionType := req.ActionType

	// 检查登录状态
	if myID == 0 {
		resp.StatusCode = 1008
		return &resp, err
	}

	if actionType == 1 {
		// 点赞,操作数据库：
		err := db.LikeAction(ctx, myID, vdID)
		if err != nil {
			resp.StatusCode = 2101
			return &resp, err
		}
	} else {
		// 取消点赞,操作数据库
		err := db.UnLikeAction(ctx, myID, vdID)
		if err != nil {
			resp.StatusCode = 2102
			return &resp, err
		}
	}

	return &resp, nil
}
