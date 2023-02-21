package userservice_test

import (
	"context"
	"testing"

	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
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
func BenchmarkUserGetFeedService(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//设置传入参数
		Token := "520222"
		req := user.FeedRequest{Token: &Token}
		userservice.UserGetFeedService(context.Background(), &req)
	}
}
func BenchmarkUserGetFeedServiceParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//设置传入参数
			Token := "520222"
			req := user.FeedRequest{Token: &Token}
			userservice.UserGetFeedService(context.Background(), &req)
		}
	})
}
