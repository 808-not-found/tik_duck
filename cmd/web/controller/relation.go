package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/808-not-found/tik_duck/cmd/web/rpc"
	"github.com/808-not-found/tik_duck/kitex_gen/useruser"
	"github.com/cloudwego/hertz/pkg/app"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

func RelationAction(ctx context.Context, c *app.RequestContext) {
	var feedReq useruser.RelationActionRequest
	feedReq.Token = c.Query("token")
	feedReq.ToUserId, _ = strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	feedReq.ActionType = int32(actionType)
	resp, err := rpc.UserRelationAction(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func FollowList(ctx context.Context, c *app.RequestContext) {
	var feedReq useruser.RelationFollowListRequest
	feedReq.Token = c.Query("token")
	feedReq.UserId, _ = strconv.ParseInt(c.Query("user_id"), 10, 64)

	resp, err := rpc.UserRelationFollowList(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func FollowerList(ctx context.Context, c *app.RequestContext) {
	var feedReq useruser.RelationFollowerListRequest
	feedReq.Token = c.Query("token")
	feedReq.UserId, _ = strconv.ParseInt(c.Query("user_id"), 10, 64)

	resp, err := rpc.UserRelationFollowerList(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func FriendList(ctx context.Context, c *app.RequestContext) {
	var feedReq useruser.RelationFriendListRequest
	feedReq.Token = c.Query("token")
	feedReq.UserId, _ = strconv.ParseInt(c.Query("user_id"), 10, 64)

	resp, err := rpc.UserRelationFriendList(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
