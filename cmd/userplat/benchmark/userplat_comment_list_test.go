package userplatservice_test

import (
	"context"
	"testing"

	userplatservice "github.com/808-not-found/tik_duck/cmd/userplat/userplatService"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
)

//	struct CommentListRequest {
//	    1: string Token (go.tag = 'json:"token"')//用户鉴权token
//	    2: i64 VideoId (go.tag = 'json:"video_id"')//视频id
//	}
//
//	struct CommentListResponse {
//	    1: i32 StatusCode (go.tag = 'json:"status_code"')//状态码，0- 成功，其他值失败
//	    2: optional string StatusMsg (go.tag = 'json:"status_msg"')//返回状态描述
//	    3: list<Comment> CommentList (go.tag = 'json:"comment_list"')//评论列表
//	}
func BenchmarkUserCommentListService(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//设置传入参数
		req := userplat.CommentListRequest{
			Token:   "3",
			VideoId: 777,
		}
		_, _ = userplatservice.UserCommentListService(context.Background(), &req)
	}
}
func BenchmarkUserCommentListServiceParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//设置传入参数
			req := userplat.CommentListRequest{
				Token:   "3",
				VideoId: 777,
			}
			_, _ = userplatservice.UserCommentListService(context.Background(), &req)
		}
	})
}
