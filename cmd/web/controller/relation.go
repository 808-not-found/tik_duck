package controller

import (
	"context"
	"log"
	"net/http"

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
	if err := c.Bind(&feedReq); err != nil {
		log.Fatalln(err)
		return
	}
	resp, err := rpc.UserRelationAction(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func FollowList(ctx context.Context, c *app.RequestContext) {
	var feedReq useruser.RelationFollowListRequest
	if err := c.Bind(&feedReq); err != nil {
		log.Fatalln(err)
		return
	}
	resp, err := rpc.UserRelationFollowList(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func FollowerList(ctx context.Context, c *app.RequestContext) {
	var feedReq useruser.RelationFollowerListRequest
	if err := c.Bind(&feedReq); err != nil {
		log.Fatalln(err)
		return
	}
	resp, err := rpc.UserRelationFollowerList(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func FriendList(ctx context.Context, c *app.RequestContext) {
	var feedReq useruser.RelationFriendListRequest
	if err := c.Bind(&feedReq); err != nil {
		log.Fatalln(err)
		return
	}
	resp, err := rpc.UserRelationFriendList(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
