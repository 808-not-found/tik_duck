namespace * useruser;

struct RelationActionRequest {
    1: string Token //用户鉴权token
    2: i64 ToUserId //对方用户id
    3: i32 ActionType // 1-关注，2-取消关注
}
struct RelationActionResponse {
    1: i32 StatusCode //状态码，0- 成功，其他值-失败
    2: string StatusMsg //返回状态描述
}
struct RelationFollowListRequest {
    1: i64 UserId //用户id
    2: string Token //用户鉴权token
}
struct RelationFollowListResponse {
    1: i32 StatusCode //状态码，0成功，其他值-失败
    2: string StatusMsg //返回状态描述
    3: list<User> UserList //用户信息列表
}
struct User {
    1: i64 Id //用户id
    2: string Name //用户名称
    3: i64 FollowCount //关注总数
    4: i64 FollowerCount //粉丝总数
    5: bool IsFollow // true- 已关注，false-未关注
}
struct RelationFollowerListRequest {
    1: i64 UserId //用户id
    2: string Token //用户鉴权token
}
struct RelationFollowerListResponse {
    1: i32 StatusCode //状态码，0- 成功，其他值失败
    2: string StatusMsg //返回状态描述
    3: list<User> UserList //用户列表
}
struct RelationFriendListRequest {
    1: i64 UserId //用户id
    2: string Token //用户鉴权token
}
struct RelationFriendListResponse {
    1: i32 StatusCode //状态码，0- 成功，其他值失败
    2: string StatusMsg //返回状态描述
    3: list<User> UserList //用户列表
}

service UserUserService {
    RelationActionResponse UserRelationAction(RelationActionRequest Req)
    RelationListResponse UserRelationList(RelationListRequest Req)
    
}