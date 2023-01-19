package controller

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

// FavoriteAction no practical effect, just check if token is valid.
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list.
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
