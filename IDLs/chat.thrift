namespace * chat

struct MessageChatRequest {
    1: string Token (go.tag = 'json:"token"')
    2: i64 ToUserId (go.tag = 'json:"to_user_id"')
    3: i64 PreMsgTime (go.tag = 'json:"pre_msg_time"')
}

struct MessageChatResponse {
    1: i32 StatusCode (go.tag = 'json:"status_code"')
    2: optional string StatusMsg (go.tag = 'json:"status_msg"')
    3: list<Message> MessageList (go.tag = 'json:"message_list"')
}

struct Message {
    1: i64 Id (go.tag = 'json:"id"')
    2: i64 ToUserId (go.tag = 'json:"to_user_id"')
    3: i64 FromUserId (go.tag = 'json:"from_user_id"')
    4: string Content (go.tag = 'json:"content')
    5: optional string CreateTime (go.tag = 'json:"create_time"')
}

struct RelationActionRequest {
    1: string Token (go.tag = 'json:"token"')
    2: i64 ToUserId (go.tag = 'json:"to_user_id"')
    3: i32 ActionType (go.tag = 'json:"action_type')
    4: string Content (go.tag = 'json:"content')
}

struct RelationActionResponse{
    1: i32 StatusCode (go.tag = 'json:"status_code"')
    2: optional string StatusMsg (go.tag = 'json:"status_msg"')
}

service ChatService {
    MessageChatResponse GetMsg (1: MessageChatRequest Req)
    RelationActionResponse PostMsg (1: RelationActionRequest Req)
}