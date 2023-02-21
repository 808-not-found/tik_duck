package userplatservice_test

import (
	"context"
	"testing"

	userplatservice "github.com/808-not-found/tik_duck/cmd/userplat/userplatService"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
)

// struct FavoriteActionRequest {
//     1: string Token (go.tag = 'json:"token"')
//     2: i64 VideoId (go.tag = 'json:"video_id"')
//     3: i32 ActionType (go.tag = 'json:"action_type"')
// }

//	struct FavoriteActionResponse {
//	    1: i32 StatusCode   (go.tag = 'json:"status_code"')        //状态码，0-成功，其他值失败
//	    2: string StatusMsg (go.tag = 'json:"status_msg"')        // 返回状态描述
//	}
func BenchmarkUserFavoriteActionService(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//设置传入参数
		req := userplat.FavoriteActionRequest{
			Token:      "3",
			VideoId:    999,
			ActionType: 1,
		}
		userplatservice.UserFavoriteActionService(context.Background(), &req)
	}
}
func BenchmarkUserFavoriteActionServiceParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//设置传入参数
			req := userplat.FavoriteActionRequest{
				Token:      "3",
				VideoId:    999,
				ActionType: 1,
			}
			userplatservice.UserFavoriteActionService(context.Background(), &req)
		}
	})
}
