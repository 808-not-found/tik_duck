package userservice_test

import (
	"context"
	"testing"

	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
)

//	struct PublishActionRequest {
//	    1: string Token (go.tag = 'json:"token"') //用户鉴权token
//	    2: string FilePath (go.tag = 'json:"file_path"')  // 视频路径
//	    3: string CoverPath (go.tag = 'json:"cover_path"') // 封面路径
//	    4: string Title (go.tag = 'json:"title"') //视频标题
//	}
//
//	struct PublishActionResponse {
//	    1: i32 StatusCode (go.tag = 'json:"status_code"') //状态码，0-成功，其他值-失败
//	    2: optional string StatusMsg (go.tag = 'json:"status_msg"') //返回状态描述
//	}
func BenchmarkUserPublishActionService(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//设置传入参数
		req := user.PublishActionRequest{
			Token:     "1926",
			FilePath:  "public/0.mp4",
			CoverPath: "public/8.jpg",
			Title:     "17",
		}
		_, _ = userservice.UserPublishActionService(context.Background(), &req)
	}
}
func BenchmarkUserPublishActionServiceParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//设置传入参数
			req := user.PublishActionRequest{
				Token:     "1926",
				FilePath:  "public/0.mp4",
				CoverPath: "public/8.jpg",
				Title:     "17",
			}
			_, _ = userservice.UserPublishActionService(context.Background(), &req)
		}
	})
}
