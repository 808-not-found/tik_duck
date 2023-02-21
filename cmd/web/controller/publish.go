package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/808-not-found/tik_duck/cmd/web/rpc"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/cover"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory.
func Publish(ctx context.Context, c *app.RequestContext) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	// 用户鉴权
	if _, err := jwt.ParseToken(token); err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	quser, _ := jwt.ParseToken(token)

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%s_%s", quser.Username, filename)
	saveFile := filepath.Join("./public/", finalName)
	saveCover := strings.ReplaceAll(saveFile, ".mp4", ".png")
	if err2 := c.SaveUploadedFile(data, saveFile); err2 != nil {
		hlog.Error(err2)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err2.Error(),
		})
		return
	}
	cover.GetSnapshot(saveFile, saveCover, 1)
	var userPublishActionReq user.PublishActionRequest
	userPublishActionReq.CoverPath = saveCover
	userPublishActionReq.FilePath = saveFile
	userPublishActionReq.Title = title
	userPublishActionReq.Token = token
	resp, err := rpc.UserPublishAction(context.Background(), &userPublishActionReq)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// PublishList all users have same publish video list.
func PublishList(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	userID := c.Query("user_id")
	var userPublishListReq user.PublishListRequest
	userPublishListReq.Token = token
	userPublishListReq.UserId, _ = strconv.ParseInt(userID, 10, 64)
	resp, err := rpc.UserPublishList(context.Background(), &userPublishListReq)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
