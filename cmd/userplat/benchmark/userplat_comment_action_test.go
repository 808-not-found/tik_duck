package userplatservice_test

import (
	"context"
	"testing"

	userplatservice "github.com/808-not-found/tik_duck/cmd/userplat/userplatService"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
)

//	struct CommentActionRequest {
//	   1: string Token (go.tag = 'json:"token"')//用户鉴权token
//	   2: i64 VideoId (go.tag = 'json:"video_id"')//视频id
//	   3: i32 ActionType (go.tag = 'json:"action_type"')// 1- 发布评论，2- 删除评论
//	   4: optional string CommentText (go.tag = 'json:"comment_text"')//用户填写的评论内容，在action_type=1 的时候使用
//	   5: optional i64 CommentId (go.tag = 'json:"comment_id"')//要删除的评论id,在action_type=2的时候使用
//	}
//
//	struct CommentActionResponse {
//	   1: i32 StatusCode (go.tag = 'json:"status_code"')//状态码，0- 成功，其他值失败
//	   2: optional string StatusMsg (go.tag = 'json:"status_msg"')//返回状态描述
//	   3: optional Comment Comment (go.tag = 'json:"comment"')//评论成功返回评论内容，不需要重新拉取整个列表
//	}
func BenchmarkUserCommentActionService(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//设置传入参数
		rettext := "I hate you"
		req := userplat.CommentActionRequest{
			Token:       "2333",
			VideoId:     999,
			ActionType:  1,
			CommentText: &rettext,
		}
		userplatservice.UserCommentActionService(context.Background(), &req)
	}
}
func BenchmarkUserCommentActionServiceParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//设置传入参数
			rettext := "I hate you"
			req := userplat.CommentActionRequest{
				Token:       "2333",
				VideoId:     999,
				ActionType:  1,
				CommentText: &rettext,
			}
			userplatservice.UserCommentActionService(context.Background(), &req)
		}
	})
}
