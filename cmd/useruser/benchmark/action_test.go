package uuservice_test

import (
	"context"
	"testing"

	uuservice "github.com/808-not-found/tik_duck/cmd/useruser/uuservice"
	useruser "github.com/808-not-found/tik_duck/kitex_gen/useruser"
	//"github.com/808-not-found/tik_duck/pkg/consts"
	//"github.com/808-not-found/tik_duck/cmd/useruser/pack"
)

//	struct RelationActionRequest {
//	    1: string Token (go.tag = 'json:"token"')//用户鉴权token
//	    2: i64 ToUserId (go.tag = 'json:"to_user_id"')//对方用户id
//	    3: i32 ActionType (go.tag = 'json:"action_type"')// 1-关注，2-取消关注
//	}
//
//	struct RelationActionResponse {
//	    1: i32 StatusCode (go.tag = 'json:"status_code"')//状态码，0- 成功，其他值-失败
//	    2: optional string StatusMsg (go.tag = 'json:"status_msg"')//返回状态描述
//	}

func BenchmarkUserRelationActionService(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//设置传入参数
		req := useruser.RelationActionRequest{
			Token:      "1231312",
			ToUserId:   1231231,
			ActionType: 1,
		}
		_, _ = uuservice.UserRelationActionService(context.Background(), &req)
	}
}
func BenchmarkUserRelationActionServiceParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//设置传入参数
			req := useruser.RelationActionRequest{
				Token:      "1231312",
				ToUserId:   1231231,
				ActionType: 1,
			}
			_, _ = uuservice.UserRelationActionService(context.Background(), &req)
		}
	})
}
