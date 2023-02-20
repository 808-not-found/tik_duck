package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/808-not-found/tik_duck/cmd/web/rpc"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
	"github.com/cloudwego/hertz/pkg/app"
)

func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var feedReq userplat.FavoriteActionRequest
	feedReq.VideoId, _ = strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	feedReq.ActionType = int32(actionType)
	feedReq.Token = c.Query("token")
	resp, err := rpc.UserFavoriteAction(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var feedReq userplat.FavoriteListRequest
	feedReq.Token = c.Query("token")
	feedReq.UserId, _ = strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err := c.Bind(&feedReq); err != nil {
		log.Fatalln(err)
		return
	}
	resp, err := rpc.UserFavoriteList(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
