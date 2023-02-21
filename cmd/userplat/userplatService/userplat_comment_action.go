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
	resp := userplat.CommentActionResponse{
		StatusCode: 0,
	}
	// 用户鉴权
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = 2034
		return &resp, nil
	}
	myID := claims.ID
	vdID := req.VideoId
	actionType := req.ActionType
	commentText := req.CommentText
	commentID := req.CommentId

	if actionType == 1 {
		// 评论
		var dbComment *db.Comment
		dbComment, err = db.CommentAction(ctx, myID, vdID, *commentText)
		if err != nil {
			resp.StatusCode = 2035
			return &resp, err
		}
		// 封装
		var rpcComment *userplat.Comment
		rpcComment, err = pack.Comment(ctx, dbComment, myID)
		if err != nil {
			resp.StatusCode = 2036
			return &resp, err
		}
		resp.Comment = rpcComment
		return &resp, nil
	}
	// 取消评论
	err = db.UnCommentAction(ctx, myID, vdID, *commentID)
	if err != nil {
		resp.StatusCode = 2035
		return &resp, err
	}

	return &resp, nil
}
