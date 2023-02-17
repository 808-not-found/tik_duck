package userplatservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/userplat/dal/db"
	"github.com/808-not-found/tik_duck/cmd/userplat/pack"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func UserCommentActionService(
	ctx context.Context,
	req *userplat.CommentActionRequest,
) (*userplat.CommentActionResponse, error) {
	var resp userplat.CommentActionResponse

	// 用户鉴权
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = 1007
		return &resp, nil
	}
	// 获取必要信息
	// 1.获取登录用户ID
	// 2.获取当前视频ID
	// 检查是进行评论还是删除
	myID := claims.ID
	vdID := req.VideoId
	actionType := req.ActionType
	commentText := req.CommentText
	commentID := req.CommentId
	// 检查登录状态
	if myID == 0 {
		resp.StatusCode = 1008
		return &resp, err
	}
	if actionType == 1 {
		// 评论,操作数据库：
		//  返回两个值一条为评论的信息，目前没有处理完这步

		var dbComments *db.Comment
		// dbComments, err = db.GetFavoriteList(ctx, req.UserId)
		dbComments, err := db.CommentAction(ctx, myID, vdID, *commentText)
		if err != nil {
			resp.StatusCode = 2101
			return &resp, err
		}
		// 数据库封装
		rpcComments, err := pack.Comment(ctx, dbComments, myID)
		if err != nil {
			resp.StatusCode = 1007
			return &resp, err
		}
		resp.Comment = rpcComments

		// return &resp, nil
	} else {
		// 取消评论,操作数据库
		err := db.UnCommentAction(ctx, myID, vdID, *commentID)
		if err != nil {
			resp.StatusCode = 2102
			return &resp, err
		}
	}

	return &resp, nil
}
