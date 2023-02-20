package userservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/cmd/user/pack"
	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/allerrors"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

//	struct PublishListRequest {
//	   1: i64 UserId (go.tag = 'json:"user_id"') //用户id
//	   2: string Token (go.tag = 'json:"token"') //用户鉴权token
//	}
//
//	struct PublishListResponse {
//	   1: i32 StatusCode (go.tag = 'json:"status_code"') //状态码，0-成功，其他值-失败
//	   2: optional string StatusMsg (go.tag = 'json:"status_msg"') //返回状态描述
//	   3: list<Video> VideoList (go.tag = 'json:"video_list"') //用户发布的视频列表
//	}
func TestUserPublishListService(t *testing.T) {
	nowTime := time.Now()
	retVideo := make([]*db.Video, 0)
	retVideo = append(retVideo, &db.Video{
		ID: 1, AuthorID: 1, PublishTime: nowTime, FilePath: "public/123.mp4", CoverPath: "public/123.jpg",
		FavoriteCount: 0, CommentCount: 0,
		Title: "test",
	},
	)
	expectVideo := make([]*user.Video, 0)
	expectVideo = append(expectVideo, &user.Video{
		Id: 1, Author: &user.User{
			Id:            1,
			Name:          "蒂萨久",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		},
		PlayPath:      "http://" + consts.WebServerPublicIP + ":" + consts.StaticPort + "/" + "public/123.mp4",
		CoverPath:     "public/123.jpg",
		FavoriteCount: 0, CommentCount: 0,
		IsFavorite: false, Title: "test"},
	)
	// 用户鉴权失败
	PatchConvey("TestMockUserPublishList_WorryClaims", t, func() {
		expectstatusCode := int32(1026)
		var expectStatuMsg *string
		var expectVideo []*user.Video
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, allerrors.ErrTestnotnil()).Build()
		req := user.PublishListRequest{
			UserId: 222,
			Token:  "777",
		}
		res, err := userservice.UserPublishListService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatuMsg, res.StatusMsg)
		assert.Equal(t, expectVideo, res.VideoList)
		assert.Equal(t, err, nil)
	})
	// 获取数据失败
	PatchConvey("TestMockUserPublishList_WorrydbVideo", t, func() {
		expectstatusCode := int32(1027)
		var expectStatuMsg *string
		var expectVideo []*user.Video
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.UserPublishList).Return(retVideo, allerrors.ErrTestnotnil()).Build()
		req := user.PublishListRequest{
			UserId: 222,
			Token:  "777",
		}
		res, err := userservice.UserPublishListService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatuMsg, res.StatusMsg)
		assert.Equal(t, expectVideo, res.VideoList)
		assert.Equal(t, err, nil)
	})
	// 封装数据失败
	PatchConvey("TestMockUserPublishList_WorryrpcVideo", t, func() {
		expectstatusCode := int32(1028)
		var expectStatuMsg *string
		var expectVideo []*user.Video
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.UserPublishList).Return(retVideo, nil).Build()
		Mock(pack.Videos).Return(expectVideo, allerrors.ErrTestnotnil()).Build()
		req := user.PublishListRequest{
			UserId: 222,
			Token:  "777",
		}
		res, err := userservice.UserPublishListService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatuMsg, res.StatusMsg)
		assert.Equal(t, expectVideo, res.VideoList)
		assert.Equal(t, err, nil)
	})
	// 正常情况
	PatchConvey("TestMockUserPublishList_normal", t, func() {
		expectstatusCode := int32(0)
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.UserPublishList).Return(retVideo, nil).Build()
		Mock(pack.Videos).Return(expectVideo, nil).Build()
		req := user.PublishListRequest{
			UserId: 222,
			Token:  "777",
		}
		res, err := userservice.UserPublishListService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectVideo, res.VideoList)
		assert.Equal(t, err, nil)
	})
}
