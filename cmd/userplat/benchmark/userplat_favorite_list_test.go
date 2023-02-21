package userplatservice_test

import (
	"context"
	"testing"

	userplatservice "github.com/808-not-found/tik_duck/cmd/userplat/userplatService"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
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
func BenchmarkUserFavoriteListService(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//设置传入参数
		req := userplat.FavoriteListRequest{
			UserId: 222777,
			Token:  "3",
		}
		userplatservice.UserFavoriteListService(context.Background(), &req)
	}
}
func BenchmarkUserFavoriteListServiceParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//设置传入参数
			req := userplat.FavoriteListRequest{
				UserId: 222777,
				Token:  "3",
			}
			userplatservice.UserFavoriteListService(context.Background(), &req)
		}
	})
}
