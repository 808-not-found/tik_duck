package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request.
func Feed(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
