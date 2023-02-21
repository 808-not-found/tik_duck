package userservice_test

import (
	"context"
	"testing"

	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
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
func BenchmarkUserPublishListService(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//设置传入参数
		req := user.PublishListRequest{
			UserId: 222,
			Token:  "777",
		}
		userservice.UserPublishListService(context.Background(), &req)
	}
}
func BenchmarkUserPublishListServiceParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//设置传入参数
			req := user.PublishListRequest{
				UserId: 222,
				Token:  "777",
			}
			userservice.UserPublishListService(context.Background(), &req)
		}
	})
}
