package userservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/cmd/user/pack"
	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/808-not-found/tik_duck/pkg/jwt"
	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

func TestUserGetFeedService(t *testing.T) {
	nowTime := time.Now()
	retVideo := make([]*db.Video, 1)
	expectVideo := make([]*user.Video, 1)
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
	PatchConvey("TestMockUserGetFeedService", t, func() {
		Mock(db.UserGetFeed).Return(retVideo, nil).Build() // mock函数
		Mock(db.GetUser).Return(retUser, nil).Build()      // mock函数
		Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()
		Mock(pack.DBUserToRPCUser).Return(&user.User{
			Id:            1,
			Name:          "蒂萨久",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		}, nil).Build()
		Token := "123412"
		res, err := userservice.UserGetFeedService(context.Background(), &user.FeedRequest{Token: &Token})
		assert.Equal(t, res.VideoList, expectVideo)
		assert.Equal(t, err, nil)
	})
}
