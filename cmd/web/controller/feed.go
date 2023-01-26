package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/808-not-found/tik_duck/cmd/web/rpc"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func Feed(ctx context.Context, c *app.RequestContext) {
	var feedReq user.FeedRequest
	if err := c.Bind(&feedReq); err != nil {
		log.Fatalln(err)
		return
	}
	resp, err := rpc.GetFeed(context.Background(), &feedReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
