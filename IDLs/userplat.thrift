namespace * userplat

struct FavoriteActionRequest {
    1: string Token (go.tag = 'json:"token"')
    2: i64 VideoId (go.tag = 'json:"video_id"')
    3: i32 ActionType (go.tag = 'json:"action_type"')
}

struct FavoriteActionResponse {
    1: i32 StatusCode   (go.tag = 'json:"status_code"')        //状态码，0-成功，其他值失败
    2: string StatusMsg (go.tag = 'json:"status_msg"')        // 返回状态描述
}

struct FavoriteListRequest {
    1: i64 UserId (go.tag = 'json:"user_id"')//用户id
    2: string Token (go.tag = 'json:"token"') //用户鉴权token
}
struct FavoriteListResponse {
    1: i32 StatusCode (go.tag = 'json:"status_code"') //状态码，0-成功，其他值失败
    2: optional string StatusMsg (go.tag = 'json:"status_msg"')//返回状态描述
    3: list<Video> VideoList (go.tag = 'json:"video_list"')//用户点赞视频列表
}
struct Video {
    1: i64 Id (go.tag = 'json:"id"')//视频唯一标识
    2: User Author (go.tag = 'json:"author"')//视频作者信息
    3: string PlayUrl (go.tag = 'json:"play_url"')//视频播放地址
    4: string CoverUrl (go.tag = 'json:"cover_url"')//视频封面地址
    5: i64 FavoriteCount (go.tag = 'json:"favorite_count"')//视频的点赞总数
    6: i64 CommentCount (go.tag = 'json:"comment_count"')//视频的评论总数
    7: bool IsFavorite (go.tag = 'json:"is_favorite"')// true- 已点赞，false-未点赞
    8: string Title (go.tag = 'json:"title"')//视频标题
}
struct User {
    1: i64 Id (go.tag = 'json:"id"')//用户id
    2: string Name (go.tag = 'json:"name"')//用户名称
    3: optional i64 FollowCount (go.tag = 'json:"follow_count"')//关注总数
    4: optional i64 FollowerCount (go.tag = 'json:"follower_count"')//粉丝总数
    5: bool IsFollow (go.tag = 'json:"is_follow"')// true-已关注，false-未关注
}
struct CommentActionRequest {
    1: string Token (go.tag = 'json:"token"')//用户鉴权token
    2: i64 VideoId (go.tag = 'json:"video_id"')//视频id
    3: i32 ActionType (go.tag = 'json:"action_type"')// 1- 发布评论，2- 删除评论
    4: optional string CommentText (go.tag = 'json:"comment_text"')//用户填写的评论内容，在action_type=1 的时候使用
    5: optional i64 CommentId (go.tag = 'json:"comment_id"')//要删除的评论id,在action_type=2的时候使用
}
struct CommentActionResponse {
    1: i32 StatusCode (go.tag = 'json:"status_code"')//状态码，0- 成功，其他值失败
    2: optional string StatusMsg (go.tag = 'json:"status_msg"')//返回状态描述
    3: optional Comment Comment (go.tag = 'json:"comment"')//评论成功返回评论内容，不需要重新拉取整个列表
}
struct Comment {
    1: i64 Id (go.tag = 'json:"id"')//视频评论id
    2: User User (go.tag = 'json:"user"')//评论用户信息
    3: string Content (go.tag = 'json:"content"')//评论内容
    4: string CreateDate (go.tag = 'json:"creaet_date"')//评论发布日期，格式mm-dd
}
struct CommentListRequest {
    1: string Token (go.tag = 'json:"token"')//用户鉴权token
    2: i64 VideoId (go.tag = 'json:"video_id"')//视频id
}
struct CommentListResponse {
    1: i32 StatusCode (go.tag = 'json:"status_code"')//状态码，0- 成功，其他值失败
    2: optional string StatusMsg (go.tag = 'json:"status_msg"')//返回状态描述
    3: list<Comment> CommentList (go.tag = 'json:"comment_list"')//评论列表
}


service UserPlatService {
    // 用户点赞
    FavoriteActionResponse UserFavoriteAction(1:FavoriteActionRequest Req)
    // 用户点赞列表
    FavoriteListResponse UserFavoriteList(1:FavoriteListRequest Req)
    // 用户评论
    CommentActionResponse UserCommentAction(1:CommentActionRequest Req)
    // 用户评论列表
    CommentListResponse UserCommentList(1:CommentListRequest Req)
}