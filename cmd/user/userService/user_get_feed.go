package userservice

import (
	"context"
	"time"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/cmd/user/pack"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/consts"
)

// statusCode, statusMsg, videoList, nextTime, err := userservice.UserGetFeedService(ctx, req).
func UserGetFeedService(ctx context.Context, req *user.FeedRequest) (int32, string, []*user.Video, int64, error) {
	// 创建返回参数
	var statusCode int32
	var statusMsg string
	var videoList []*user.Video
	var nextTime int64
	var latestTime time.Time
	// 完善用户请求
	if req.LatestTime != nil {
		latestTime = time.Unix(*req.LatestTime, 0)
	} else {
		latestTime = time.Now()
	}

	// 向数据库获取视频列表
	dbVideoList, nextTime, err := db.UserGetFeed(ctx, latestTime)
	if err != nil {
		statusCode = 1102
		return statusCode, statusMsg, nil, nextTime, err
	}

	// 将数据库数据转化为接口定义的结构体并打包
	videoList = pack.Videos(dbVideoList)

	// 成功返回
	statusMsg = consts.Success
	return statusCode, statusMsg, videoList, nextTime, nil
}
