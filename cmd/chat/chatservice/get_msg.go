package chatservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/chat/dal/db"
	"github.com/808-not-found/tik_duck/cmd/chat/pack"
	chat "github.com/808-not-found/tik_duck/kitex_gen/chat"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func GetMsgService(ctx context.Context, req *chat.MessageChatRequest) (resp *chat.MessageChatResponse, err error) {
	resp = &chat.MessageChatResponse{
		StatusCode: 0,
	}

	var myID int64
	if req.Token == "" {
		resp.StatusCode = 4001
		return resp, err
	}
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = 4002
		return resp, err
	}
	myID = claims.ID

	var chatlist []*db.Message
	chatlist, err = db.GetMsg(ctx, myID, req.ToUserId, &req.PreMsgTime)
	if err != nil {
		resp.StatusCode = 4003
		return resp, err
	}

	var rpcList []*chat.Message
	rpcList, err = pack.Msgs(ctx, chatlist)
	if err != nil {
		resp.StatusCode = 4004
		return resp, err
	}

	resp = &chat.MessageChatResponse{
		StatusCode: 0,
		MessageList: rpcList,
	}
	return resp, nil
}