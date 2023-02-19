package userservice

import (
	"context"
	"time"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/cmd/user/pack"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

// statusCode, statusMsg, videoList, nextTime, err := userservice.UserGetFeedService(ctx, req).
func UserGetFeedService(ctx context.Context, req *user.FeedRequest) (*user.FeedResponse, error) {
	var msg string
	resp := user.FeedResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
	}

	// 验证登录状态
	var myID int64
	if req.Token == nil {
		myID = 0
	} else {
		claims, err := jwt.ParseToken(*req.Token)
		if err != nil {
			resp.StatusCode = 1021
			return &resp, err
		}
		myID = claims.ID
	}

	// 查询数据库 videolist
	dbVideos, err := db.UserGetFeed(ctx, req.LatestTime)
	if err != nil {
		resp.StatusCode = 1022
		return &resp, err
	}

	// 获取最早时间
	var nextTime int64
	if len(dbVideos) == 0 {
		nextTime = time.Now().Unix()
	} else {
		nextTime = dbVideos[len(dbVideos)-1].PublishTime.Unix()
	}

	// 封装数据
	rpcVideos, err := pack.Videos(ctx, dbVideos, myID)
	if err != nil {
		resp.StatusCode = 1023
		return &resp, err
	}
	res := user.FeedResponse{
		StatusCode: 0,
		VideoList:  rpcVideos,
		NextTime:   &nextTime,
	}

	// 返回数据
	return &res, err
}
