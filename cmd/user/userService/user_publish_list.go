package userservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/cmd/user/pack"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/consts"
)

func UserPublishListService(ctx context.Context, userID int32) (int32, string, []*user.Video, error) {
	// 创建返回参数
	var statusCode int32
	var statusMsg string
	var videoList []*user.Video

	// 向数据库获取视频列表
	dbvideoList, err := db.UserPublishList(ctx, userID)
	if err != nil {
		statusCode = 1102
		return statusCode, statusMsg, nil, err
	}
	// 将数据打包
	videoList = pack.Videos(dbvideoList)

	// 成功返回
	statusMsg = consts.Success
	return statusCode, statusMsg, videoList, nil
}
