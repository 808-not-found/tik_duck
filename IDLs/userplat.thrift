namespace * userplat;

struct DouyinFavoriteListRequest {
    1: i64 UserId //用户id
    2: string Token //用户鉴权token
}
struct DouyinFavoriteListResponse {
    1: i32 StatusCode //状态码，0-成功，其他值失败
    2: string StatusMsg //返回状态描述
    3: list<Video> VideoList //用户点赞视频列表
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
struct DouyinCommentActionRequest {
    1: string Token //用户鉴权token
    2: i64 VideoId //视频id
    3: i32 ActionType // 1- 发布评论，2- 删除评论
    4: string CommentText //用户填写的评论内容，在action_type=1 的时候使用
    5: i64 CommentId //要删除的评论id,在action_type=2的时候使用
}
struct DouyinCommentActionResponse {
    1: i32 StatusCode //状态码，0- 成功，其他值失败
    2: string StatusMsg //返回状态描述
    3: Comment Comment //评论成功返回评论内容，不需要重新拉取整个列表
}
struct Comment {
    1: i64 Id //视频评论id
    2: User User //评论用户信息
    3: string Content //评论内容
    4: string CreateDate //评论发布日期，格式mm-dd
}
struct DouyinCommentListRequest {
    1: string Token //用户鉴权token
    2: i64 VideoId //视频id
}
struct DouyinCommentListResponse {
    1: i32 StatusCode //状态码，0- 成功，其他值失败
    2: string StatusMsg //返回状态描述
    3: list<Comment> CommentList //评论列表
}

service UserPlatService {
    i32 sayInt(1:i32 param)
    string sayString(1:string param)
    bool sayBoolean(1:bool param)
    void sayVoid()
}