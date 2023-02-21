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

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin.
//
//nolint:all
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		ID:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

//nolint:all
var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserID int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(ctx context.Context, c *app.RequestContext) {
	var userRegisterReq user.UserRegisterRequest
	userRegisterReq.Username = c.Query("username")
	userRegisterReq.Password = c.Query("password")
	resp, err := rpc.UserRegister(context.Background(), &userRegisterReq)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, err)
	}
	c.JSON(http.StatusOK, resp)
}

func Login(ctx context.Context, c *app.RequestContext) {
	var userLoginReq user.UserLoginRequest
	userLoginReq.Username = c.Query("username")
	userLoginReq.Password = c.Query("password")

	resp, err := rpc.UserLogin(context.Background(), &userLoginReq)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, err)
	}
	c.JSON(http.StatusOK, resp)
}

func UserInfo(ctx context.Context, c *app.RequestContext) {
	var userInfoReq user.UserRequest
	userInfoReq.Token = c.Query("token")
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, err)
	}
	userInfoReq.UserId = userID
	resp, err := rpc.UserInfo(context.Background(), &userInfoReq)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, err)
	}
	c.JSON(http.StatusOK, resp)
}
