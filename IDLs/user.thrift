namespace * user;

struct DouyinFeedRequest {
    1: i64 LatestTime //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: string Token // 可选参数，登录用户设置
}
struct DouyinFeedResponse {
    1: i32 StatusCode //状态码，0-成功，其他值失败
    2: string StatusMsg //返回状态描述
    3: list<Video> VideoList //视频列表
    4: i64 NextTime //本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}
struct Video {
    1: i64 Id //视频唯一标识
    2: User Author //视频作者信息
    3: string PlayUrl //视频播放地址
    4: string CoverUrl //视频封面地址
    5: i64 FavoriteCount //视频的点赞总数
    6: i64 CommentCount //视频的评论总数
    7: bool IsFavorite // true- 已点赞，false-未点赞
    8: string Title //视频标题
}
struct User {
    1: i64 Id //用户id
    2: string Name //用户名称
    3: i64 FollowCount //关注总数
    4: i64 FollowerCount //粉丝总数
    5: bool IsFollow // true-已关注，false-未关注
}
struct DouyinUserRegisterRequest {
    1: string Username //注册用户名，最长32个字符1525
    2: string Password //密码，最长32个字符
}
struct DouyinUserRegisterResponse {
    1: i32 StatusCode //状态码，0-成功，其他值失败
    2: string StatusMsg //返回状态描述
    3: i64 UserId //用户id
    4: string Token //用户鉴权token
}
struct DouyinUserLoginRequest {
    1: string Username //登录用户名
    2: string Password //登录密码
}
struct DouyinUserLoginResponse {
    1: i32 StatusCode //状态码，0-成功，其他值失败
    2: string StatusMsg //返回状态描述
    3: i64 UserId //用户id
    4: string Token //用户鉴权token
}
struct DouyinUserRequest {
    1: i64 UserId //用户id
    2: string Token //用户鉴权token
}
struct DouyinUserResponse {
    1: i32 StatusCode //状态码，0-成功，其他值-失败
    2: string StatusMsg //返回状态描述
    3: User User //用户信息
}
struct DouyinPublishActionRequest {
    1: string Token //用户鉴权token
    2: Bytes Data //视频数据
    3: string Title //视频标题
}
struct DouyinPublishActionResponse {
    1: i32 StatusCode //状态码，0-成功，其他值-失败
    2: string StatusMsg //返回状态描述
}
struct DouyinPublishListRequest {
    1: i64 UserId //用户id
    2: string Token //用户鉴权token
}
struct DouyinPublishListResponse {
    1: i32 StatusCode //状态码，0-成功，其他值-失败
    2: string StatusMsg //返回状态描述
    3: list<Video> VideoList //用户发布的视频列表
}

service UserService {
    //注册
    DouyinUserRegisterResponse UserRegister (1: DouyinUserRegisterRequest Req)
    //获取视频流
    DouyinFeedResponse UserGetFeed (1: DouyinFeedRequest Req)
    //登录
    DouyinUserLoginResponse UserLogin (1: DouyinUserLoginRequest Req)
    //获取用户信息
    DouyinUserInfoResponse UserInfo (1: DouyinUserInfoRequest Req)
    //获取用户发布作品
    DouyinPublishListResponse UserPublishList (1: DouyinPublishListRequest Req)
    //视频投稿
    DouyinPublishActionResponse UserPublishAction (1: DouyinPublishActionRequest Req)
}