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
		resp.StatusCode = 1001
		return &resp, nil
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
		resp.StatusCode = 1002
		return &resp, err
	}

	if actionType == 1 {
		// 关注
		// 操作数据库：
		// 1.增加一个关注数
		// 2.增加一个粉丝数
		// 3.向Follow表中增加一条记录
		err := db.FollowAction(myID, toID)
		if err != nil {
			resp.StatusCode = 1003
			return &resp, err
		}
	} else {
		// 取关
		// 操作数据库：
		// 1.减少一个关注数
		// 2.减少一个粉丝数
		// 3.向Follow表中删除一条记录
		err := db.UnFollowAction(myID, toID)
		if err != nil {
			resp.StatusCode = 1004
			return &resp, err
		}
	}

	return &resp, nil
}