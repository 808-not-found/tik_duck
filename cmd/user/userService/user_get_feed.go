package userservice

import (
	"context"
	"time"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/cmd/user/pack"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
)

func NewUserGetFeedService(ctx context.Context) context.Context {
	return ctx
}

// 更改了video之后 有一点小问题 数据库里面的video中 PublishTime 是 time类型
// 而你原来的是int64类型 我先改成了time类型 先过审 你回头看看是要哪个
// 视频列表.
// 学长的数据库 model里面是 time 类型 user.thrift是 int64类型

func UserGetFeedService(ctx context.Context, req *user.FeedRequest) (int32, string, []*user.Video, time.Time, error) {
	var statusCode int32
	t := time.Unix(*req.LatestTime, 0)
	videoModels, nextTime, err := db.UserGetFeed(ctx, t, *req.Token)
	if err != nil {
		statusCode = 1001
		nilTime := time.Time{}
		return statusCode, "", nil, nilTime, nil
	}

	// 成功返回

	statusMsg := "success"
	videoList := pack.Videos(videoModels)
	return statusCode, statusMsg, videoList, nextTime, nil
}

// func UserGetFeedService(ctx context.Context, req *user.FeedRequest) (int32, string, []*user.Video, int64, error) {
// 	var statusCode int32
// 	videoModels, nextTime, err := db.UserGetFeed(ctx, *req.LatestTime, *req.Token)
// 	if err != nil {
// 		statusCode = 1001
// 		return statusCode, "", nil, 0, nil
// 	}

// 	// 成功返回

//		statusMsg := "success"
//		videoList := pack.Videos(videoModels)
//		return statusCode, statusMsg, videoList, nextTime, nil
//	}
