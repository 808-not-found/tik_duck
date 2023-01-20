package userservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/cmd/user/pack"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
)

func NewUserGetFeedService(ctx context.Context) context.Context {
	return ctx
}

// 视频列表.
func UserGetFeedService(ctx context.Context, req *user.FeedRequest) (int32, string, []*user.Video, int64, error) {
	var statusCode int32
	videoModels, nextTime, err := db.UserGetFeed(ctx, *req.LatestTime, *req.Token)
	if err != nil {
		statusCode = 1001
		return statusCode, "", nil, 0, nil
	}

	// 成功返回

	statusMsg := "success"
	videoList := pack.Videos(videoModels)
	return statusCode, statusMsg, videoList, nextTime, nil
}
