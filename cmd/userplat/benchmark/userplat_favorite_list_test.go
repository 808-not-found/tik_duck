package userplatservice_test

import (
	"context"
	"testing"

	"github.com/808-not-found/tik_duck/cmd/userplat/dal/db"
	"github.com/808-not-found/tik_duck/cmd/userplat/pack"
	userplatservice "github.com/808-not-found/tik_duck/cmd/userplat/userplatService"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
	"github.com/808-not-found/tik_duck/pkg/allerrors"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

//	struct FavoriteListRequest {
//	    1: i64 UserId (go.tag = 'json:"user_id"')//用户id
//	    2: string Token (go.tag = 'json:"token"') //用户鉴权token
//	}
//
//	struct FavoriteListResponse {
//	    1: i32 StatusCode (go.tag = 'json:"status_code"') //状态码，0-成功，其他值失败
//	    2: optional string StatusMsg (go.tag = 'json:"status_msg"')//返回状态描述
//	    3: list<Video> VideoList (go.tag = 'json:"video_list"')//用户点赞视频列表
//	}
func TestUserFavoriteListService(t *testing.T) {
	//鉴权失败
	PatchConvey("TestMockUserplatFavoriteAction_WorryClaim", t, func() {
		expectstatusCode := int32(2037)
		var expectStatusMsg *string
		var expectvideo []*userplat.Video
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, allerrors.ErrTestnotnil()).Build()
		req := userplat.FavoriteListRequest{
			UserId: 222777,
			Token:  "3",
		}
		res, err := userplatservice.UserFavoriteListService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, expectvideo, res.VideoList)
		assert.Equal(t, err, nil)
	})
	//查询数据库失败
	PatchConvey("TestMockUserplatFavoriteAction_Worrydb", t, func() {
		expectstatusCode := int32(2038)
		var expectStatusMsg *string
		var expectvideo []*userplat.Video
		var retvideo []*db.Video
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.GetFavoriteList).Return(retvideo, allerrors.ErrTestnotnil()).Build()
		req := userplat.FavoriteListRequest{
			UserId: 222777,
			Token:  "3",
		}
		res, err := userplatservice.UserFavoriteListService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, expectvideo, res.VideoList)
		assert.Equal(t, err, allerrors.ErrTestnotnil())
	})
	//数据封装失败
	PatchConvey("TestMockUserplatFavoriteAction_Worryrpc", t, func() {
		expectstatusCode := int32(2039)
		var expectStatusMsg *string
		var expectvideo []*userplat.Video
		var retvideo []*db.Video
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.GetFavoriteList).Return(retvideo, nil).Build()
		Mock(pack.Videos).Return(expectvideo, allerrors.ErrTestnotnil()).Build()
		req := userplat.FavoriteListRequest{
			UserId: 222777,
			Token:  "3",
		}
		res, err := userplatservice.UserFavoriteListService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, expectvideo, res.VideoList)
		assert.Equal(t, err, allerrors.ErrTestnotnil())
	})
	//正常情况
	PatchConvey("TestMockUserplatFavoriteAction_Normal", t, func() {
		expectstatusCode := int32(0)
		var expectStatusMsg *string
		var retvideo []*db.Video
		expectVideo := make([]*userplat.Video, 0)
		expectVideo = append(expectVideo, &userplat.Video{
			Id: 1, Author: &userplat.User{
				Id:            1,
				Name:          "蒂萨久",
				FollowCount:   nil,
				FollowerCount: nil,
				IsFollow:      false,
			},
			PlayUrl:       "http://" + consts.WebServerPublicIP + ":" + consts.StaticPort + "/" + "public/123.mp4",
			CoverUrl:      "public/123.jpg",
			FavoriteCount: 0, CommentCount: 0,
			IsFavorite: false, Title: "test"},
		)
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, nil).Build()
		Mock(db.GetFavoriteList).Return(retvideo, nil).Build()
		Mock(pack.Videos).Return(expectVideo, nil).Build()
		req := userplat.FavoriteListRequest{
			UserId: 222777,
			Token:  "3",
		}
		res, err := userplatservice.UserFavoriteListService(context.Background(), &req)
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectStatusMsg, res.StatusMsg)
		assert.Equal(t, expectVideo, res.VideoList)
		assert.Equal(t, err, nil)
	})
}
