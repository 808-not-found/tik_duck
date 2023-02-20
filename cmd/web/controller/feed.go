package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"

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
	token := c.Query("token")
	latestTime, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if token != "" {
		feedReq.Token = &token
	} else {
		feedReq.Token = nil
	}
	feedReq.LatestTime = &latestTime
	resp, err := rpc.GetFeed(context.Background(), &feedReq)
	// test video
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
