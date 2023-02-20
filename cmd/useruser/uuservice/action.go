package uuservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/useruser/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/useruser"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func UserRelationActionService(
	ctx context.Context,
	req *useruser.RelationActionRequest,
) (*useruser.RelationActionResponse, error) {
	var resp useruser.RelationActionResponse

	// 用户鉴权
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = 3001
		return &resp, err
	}

	// 获取必要信息
	// 1.获取登录用户ID
	// 2.获取对方用户ID
	// 检查是关注还是取关
	myID := claims.ID
	toID := req.ToUserId
	actionType := req.ActionType

	// 检查登录状态
	if myID == 0 {
		resp.StatusCode = 3002
		return &resp, err
	}

	if actionType == 1 {
		// 关注 操作数据库
		err := db.FollowAction(ctx, myID, toID)
		if err != nil {
			resp.StatusCode = 3003
			return &resp, err
		}
	} else {
		// 取关 操作数据库
		err := db.UnFollowAction(ctx, myID, toID)
		if err != nil {
			resp.StatusCode = 3004
			return &resp, err
		}
	}

	return &resp, nil
}
