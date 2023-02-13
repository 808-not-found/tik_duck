package userplatservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/userplat/dal/db"
	"github.com/808-not-found/tik_duck/cmd/userplat/pack"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func UserCommentListService(
	ctx context.Context,
	req *userplat.CommentListRequest,
) (*userplat.CommentListResponse, error) { 
		var resp userplat.CommentListResponse

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
		
		vdID := req.VideoId
		// 查询数据库
		var dbComments []*db.Comment
		dbComments, err = db.GetCommentList(ctx, myID, vdID)
		if err != nil {
			resp.StatusCode = 1006
			return &resp, err
		}
		// 数据封装
		rpcComments, err := pack.Comments(dbComments, myID, vdID)
		if err != nil {
			resp.StatusCode = 1007
			return &resp, err
		}
		resp.CommentList = rpcComments

		return &resp, nil

}
