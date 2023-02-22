package chatservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/chat/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/chat"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func PostMsgService(
	ctx context.Context,
	req *chat.RelationActionRequest,
) (resp *chat.RelationActionResponse, err error) {
	resp = &chat.RelationActionResponse{
		StatusCode: 0,
	}
	var myID int64
	if req.Token == "" {
		resp.StatusCode = 4005
		return resp, err
	}
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = 4006
		return resp, err
	}
	myID = claims.ID

	err = db.PostMsg(ctx, myID, req.ToUserId, &req.Content)
	if err != nil {
		resp.StatusCode = 4007
		return resp, err
	}

	return resp, nil
}
