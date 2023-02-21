# 测试部分

### 1. 单元标准格式

```go
package userservice_test

import (//添加依赖
)

//  req 和 resp 格式 : 在idls中获取 可以贴在这方便观察

func TestUserGetFeedService(t *testing.T) {
    // 构建通用信息

    // 正确情况测试
    PatchConvey("TestMockXXX", t, func() {
        //设置期待值
        expectstatusCode:=int32(0)
        ...

        // 设定mock函数
        // 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
        Mock(db.UserGetFeed).Return(retVideo, nil).Build()
        ...

        //设置传入参数
        Token:="123412"

        //调用函数
        res, err:= userservice.UserGetFeedService(context.Background(), &user.FeedRequest{Token: &Token})

        //对比返回值
        assert.Equal(t, expectstatusCode, res.StatusCode)
        ...
    })

    //第二组...
    PatchConvey("TestMockXXX", t, func() {
        ...
    })

    //第三组...
}

```

例子如下：

```go
package userservice_test

import (//添加依赖
    "context"
    "testing"
    "time"
    "github.com/808-not-found/tik_duck/cmd/user/dal/db"
    "github.com/808-not-found/tik_duck/cmd/user/pack"
    userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
    "github.com/808-not-found/tik_duck/kitex_gen/user"
    "github.com/808-not-found/tik_duck/pkg/consts"
    "github.com/808-not-found/tik_duck/pkg/jwt"
    . "github.com/bytedance/mockey"
    "github.com/stretchr/testify/assert"
    "github.com/808-not-found/tik_duck/pkg/allerrors"
)

//  req 和 resp 格式 : 在idls中获取
//  struct FeedRequest {
//      1: optional i64 LatestTime (go.tag = 'json:"latest_time"') //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
//      2: optional string Token (go.tag = 'json:"token"') // 可选参数，登录用户设置
//  }
//
//  struct FeedResponse {
//      1: i32 StatusCode (go.tag = 'json:"status_code"') //状态码，0-成功，其他值失败
//      2: optional string StatusMsg (go.tag = 'json:"status_msg"') //返回状态描述
//      3: list`<Video>` VideoList (go.tag = 'json:"video_list"') //视频列表
//      4: optional i64 NextTime (go.tag = 'json:"next_time"') //本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
//  }

funcTestUserGetFeedService(t *testing.T) {
    // 构建通用信息
    nowTime:= time.Now()
    retVideo:=make([]*db.Video, 0)
    retVideo=append(retVideo, &db.Video{
        ID: 1, AuthorID: 1, PublishTime: nowTime, FilePath: "public/123.mp4", CoverPath: "public/123.jpg",
        FavoriteCount: 0, CommentCount: 0,
        Title: "test",
        },
    )
    retUser:= db.User{
        ID:            1,
        CreateTime:    nowTime,
        Name:          "蒂萨久",
        FollowCount:   0,
        FollowerCount: 0,
        Password:      "114514",
        Salt:          "1919810",
    }

    // 正确情况测试
    PatchConvey("TestMockUserGetFeedService_normal", t, func() {
        //设置期待值
        expectstatusCode:=int32(0)
        expectVideo:=make([]*user.Video, 0)
        expectVideo=append(expectVideo, &user.Video{
            Id: 1, Author: &user.User{
            Id:            1,
            Name:          "蒂萨久",
            FollowCount:   nil,
            FollowerCount: nil,
            IsFollow:      false,
            },
            PlayPath:      "http://"+ consts.WebServerPublicIP +":"+ consts.StaticPort +"/"+"public/123.mp4",
            CoverPath:     "public/123.jpg",
            FavoriteCount: 0, CommentCount: 0,
            IsFavorite: false, Title: "test"},
        )

        // 设定mock函数
        // 这部分主要是设定 被测试函数内部调用的别的函数 修改他们返回的结果
        Mock(db.UserGetFeed).Return(retVideo, nil).Build()
        Mock(db.GetUser).Return(retUser, nil).Build()
        Mock(jwt.ParseToken).Return(&jwt.MyClaims{ID: 1}, nil).Build()
        Mock(pack.DBUserToRPCUser).Return(&user.User{
            Id:            1,
            Name:          "蒂萨久",
            FollowCount:   nil,
            FollowerCount: nil,
            IsFollow:      false,
        }, nil).Build()
        //设置传入参数
        Token:="123412"

        //调用函数
        res, err:= userservice.UserGetFeedService(context.Background(), &user.FeedRequest{Token: &Token})

        //对比返回值
            assert.Equal(t, expectstatusCode, res.StatusCode)
            assert.Equal(t, expectVideo, res.VideoList)
            assert.Equal(t, err, nil)
    })
    //第二组...
    //第三组...
}

```

### 2. 基准测试格式

```go
func BenchmarkUserInfoService(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//设置传入参数
		req := user.UserRequest{
			xxxx
		}
		userservice.UserInfoService(context.Background(), &req)
	}
}
func BenchmarkUserInfoServiceParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//设置传入参数
			req := user.UserRequest{
				xxxx
			}
			userservice.UserInfoService(context.Background(), &req)
		}
	})
}

```

### 3. 测试时发现的问题

1. 关于链式方法：

   - 直接 mock 上层函数
2. 关于错误码：

   - 在 pkg/allerrors 中有错误码的具体说明
   - 尽量到达所有错误码情况以提高覆盖率
3. 当要求返回错误 err 不是 nil 时

   - 引入"github.com/808-not-found/tik_duck/pkg/allerrors"包
   - 之后使用 allerrors.ErrTestnotnil()
4. 测试命令

   - go test -gcflags="all=-l -N" -v ./...
