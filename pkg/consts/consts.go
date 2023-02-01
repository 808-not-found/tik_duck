// Some consts for database

package consts

const (
	CommentTableName    = "comment"
	FollowTableName     = "follow"
	LikeTableName       = "like"
	MessageTableName    = "message"
	UserTableName       = "user"
	VideoTableName      = "video"
	EtcdAddress         = "127.0.0.1:2379"
	UserServiceName     = "user"
	UserUserServiceName = "useruser"
	UserPlatServiceName = "userplat"

	MySQLDefaultDSN = "tik_duck:duck@tcp(localhost:9910)/tik_duck"
	TCP             = "tcp"
	Success         = "success"
	HTTPMaxBodySize = 1024 * 1024 * 1024 // 1GB
)
