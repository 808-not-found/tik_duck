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

//	struct FeedRequest {
//	    1: optional i64 LatestTime (go.tag = 'json:"latest_time"') //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
//	    2: optional string Token (go.tag = 'json:"token"') // 可选参数，登录用户设置
//	}
//
//	struct FeedResponse {
//	    1: i32 StatusCode (go.tag = 'json:"status_code"') //状态码，0-成功，其他值失败
//	    2: optional string StatusMsg (go.tag = 'json:"status_msg"') //返回状态描述
//	    3: list<Video> VideoList (go.tag = 'json:"video_list"') //视频列表
//	    4: optional i64 NextTime (go.tag = 'json:"next_time"') //本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
//	}
func TestUserGetFeedService(t *testing.T) {
	// 构建通用信息
	nowTime := time.Now()
	retVideo := make([]*db.Video, 0)
	retVideo = append(retVideo, &db.Video{
		ID: 1, AuthorID: 1, PublishTime: nowTime, FilePath: "public/123.mp4", CoverPath: "public/123.jpg",
		FavoriteCount: 0, CommentCount: 0,
		Title: "test",
	},
	)
	retUser := db.User{
		ID:            1,
		CreateTime:    nowTime,
		Name:          "蒂萨久",
		FollowCount:   0,
		FollowerCount: 0,
		Password:      "114514",
		Salt:          "1919810",
	}

	// 正确情况测试_登陆状态
	PatchConvey("TestMockUserGetFeedService_normal_WithToken", t, func() {
		//设置期待值
		expectstatusCode := int32(0)
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
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(db.UserGetFeed).Return(retVideo, nil).Build()
		Mock(db.GetUser).Return(retUser, nil).Build()
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()
		Mock(db.IsFavorite).Return(allerrors.ErrTestnotnil()).Build()
		Mock(pack.DBUserToRPCUser).Return(&user.User{
			Id:            1,
			Name:          "蒂萨久",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		}, nil).Build()

		//设置传入参数
		Token := "123412"

		//调用函数
		res, err := userservice.UserGetFeedService(context.Background(), &user.FeedRequest{Token: &Token})

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectVideo, res.VideoList)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, nil, err)
	})

	// 正确情况测试_登陆状态_关注状态
	PatchConvey("TestMockUserGetFeedService_normal_WithToken", t, func() {
		//设置期待值
		expectstatusCode := int32(0)
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
			IsFavorite: true, Title: "test"},
		)
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(db.UserGetFeed).Return(retVideo, nil).Build()
		Mock(db.GetUser).Return(retUser, nil).Build()
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()
		Mock(db.IsFavorite).Return(nil).Build()
		Mock(pack.DBUserToRPCUser).Return(&user.User{
			Id:            1,
			Name:          "蒂萨久",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		}, nil).Build()

		//设置传入参数
		Token := "123412"

		//调用函数
		res, err := userservice.UserGetFeedService(context.Background(), &user.FeedRequest{Token: &Token})

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectVideo, res.VideoList)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, nil, err)
	})

	// 正确情况测试_未登陆状态
	PatchConvey("TestMockUserGetFeedService_normal_WithoutToken", t, func() {
		//设置期待值
		expectstatusCode := int32(0)
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
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(db.UserGetFeed).Return(retVideo, nil).Build()
		Mock(db.GetUser).Return(retUser, nil).Build()
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()
		Mock(db.IsFavorite).Return(allerrors.ErrTestnotnil()).Build()
		Mock(pack.User).Return(&user.User{
			Id:            1,
			Name:          "蒂萨久",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		}).Build()

		//设置传入参数

		//调用函数
		res, err := userservice.UserGetFeedService(context.Background(), &user.FeedRequest{})

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectVideo, res.VideoList)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, nil, err)
	})

	// 验证Token失败
	PatchConvey("TestMockUserGetFeedService_WrongToken", t, func() {
		//设置期待值
		expectstatusCode := int32(1021)
		expectErr := allerrors.ErrTestnotnil()
		var expectNexttime *int64
		var expectVideo []*user.Video
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{}, allerrors.ErrTestnotnil()).Build()

		//设置传入参数
		Token := "123412"

		//调用函数
		res, err := userservice.UserGetFeedService(context.Background(), &user.FeedRequest{Token: &Token})

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectErr, err)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, expectVideo, res.VideoList)
		assert.Equal(t, expectNexttime, res.NextTime)
	})

	//查询数据库问题
	PatchConvey("TestMockUserGetFeedService_WrongVideolist", t, func() {
		//设置期待值
		expectstatusCode := int32(1022)
		expectErr := allerrors.ErrTestnotnil()
		var expectNexttime *int64
		var expectVideo []*user.Video
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(db.UserGetFeed).Return([]*db.Video{}, allerrors.ErrTestnotnil()).Build()
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()

		//设置传入参数

		//调用函数
		res, err := userservice.UserGetFeedService(context.Background(), &user.FeedRequest{})

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectErr, err)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, expectVideo, res.VideoList)
		assert.Equal(t, expectNexttime, res.NextTime)
	})

	// 封装数据错误
	PatchConvey("TestMockUserGetFeedService_WrongPackVideos", t, func() {
		//设置期待值
		expectstatusCode := int32(1023)
		expectErr := allerrors.ErrTestnotnil()
		var expectNexttime *int64
		var expectVideo []*user.Video
		var expectMsg *string

		// 设定mock函数
		// 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
		Mock(db.UserGetFeed).Return(retVideo, nil).Build()
		Mock(db.GetUser).Return(retUser, nil).Build()
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()
		Mock(pack.Videos).Return(nil, allerrors.ErrTestnotnil()).Build()
		Mock(pack.DBUserToRPCUser).Return(&user.User{
			Id:            1,
			Name:          "蒂萨久",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		}, nil).Build()

		//设置传入参数
		Token := "123412"

		//调用函数
		res, err := userservice.UserGetFeedService(context.Background(), &user.FeedRequest{Token: &Token})

		//对比返回值
		assert.Equal(t, expectstatusCode, res.StatusCode)
		assert.Equal(t, expectErr, err)
		assert.Equal(t, expectMsg, res.StatusMsg)
		assert.Equal(t, expectVideo, res.VideoList)
		assert.Equal(t, expectNexttime, res.NextTime)
	})
	//缺少时间处理
}
