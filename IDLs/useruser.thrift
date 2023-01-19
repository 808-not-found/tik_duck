namespace * useruser;

struct RelationActionRequest {
    1: string Token //用户鉴权token
    2: i64 ToUserId //对方用户id
    3: i32 ActionType // 1-关注，2-取消关注
}
struct RelationActionResponse {
    1: i32 StatusCode //状态码，0- 成功，其他值-失败
    2: optional string StatusMsg //返回状态描述
}
struct RelationFollowListRequest {
    1: i64 UserId //用户id
    2: string Token //用户鉴权token
}
struct RelationFollowListResponse {
    1: i32 StatusCode //状态码，0成功，其他值-失败
    2: optional string StatusMsg //返回状态描述
    3: list<User> UserList //用户信息列表
}
struct User {
    1: i64 Id //用户id
    2: string Name //用户名称
    3: optional i64 FollowCount //关注总数
    4: optional i64 FollowerCount //粉丝总数
    5: bool IsFollow // true- 已关注，false-未关注
}
struct RelationFollowerListRequest {
    1: i64 UserId //用户id
    2: string Token //用户鉴权token
}
struct RelationFollowerListResponse {
    1: i32 StatusCode //状态码，0- 成功，其他值失败
    2: optional string StatusMsg //返回状态描述
    3: list<User> UserList //用户列表
}
struct RelationFriendListRequest {
    1: i64 UserId //用户id
    2: string Token //用户鉴权token
}
struct RelationFriendListResponse {
    1: i32 StatusCode //状态码，0- 成功，其他值失败
    2: optional string StatusMsg //返回状态描述
    3: list<User> UserList //用户列表
}

service UserUserService {
    // 关系操作
    RelationActionResponse UserRelationAction(1:RelationActionRequest Req)
    // 关注列表
    RelationFollowListResponse UserRelationFollowList(1:RelationFollowListRequest Req)
    // 粉丝列表
    RelationFollowerListResponse UserRelationFollowerList(1:RelationFollowerListRequest Req)
    // 好友列表
    RelationFriendListResponse UserRelationFriendList(1:RelationFriendListRequest Req)
    // 消息操作暂且不实现
}