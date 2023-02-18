package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/808-not-found/tik_duck/cmd/web/rpc"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
	"github.com/cloudwego/hertz/pkg/app"
)

func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var feedReq userplat.FavoriteActionRequest
	if err := c.Bind(&feedReq); err != nil {
		log.Fatalln(err)
		return
	}
	resp, err := rpc.UserFavoriteAction(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var feedReq userplat.FavoriteListRequest
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
