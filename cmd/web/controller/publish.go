package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/808-not-found/tik_duck/cmd/web/rpc"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
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
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		hlog.Error(err)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	var userPublishActionReq user.PublishActionRequest
	userPublishActionReq.CoverPath = "http://www.hrbust.edu.cn/images/xjgk.jpg"
	userPublishActionReq.FilePath = saveFile
	userPublishActionReq.Title = title
	userPublishActionReq.Token = token
	resp, err := rpc.UserPublishAction(context.Background(), &userPublishActionReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// PublishList all users have same publish video list.
func PublishList(ctx context.Context, c *app.RequestContext) {
	var userPublishListReq user.PublishListRequest
	if err := c.Bind(&userPublishListReq); err != nil {
		log.Fatalln(err)
		return
	}
	resp, err := rpc.UserPublishList(context.Background(), &userPublishListReq)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
