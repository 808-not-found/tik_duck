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
	UserServiceName     = "usersvr"
	UserUserServiceName = "useruser"
	UserPlatServiceName = "userplat"
	ChatService         = "chat"

	WebServerPort     = "8080"
	WebServerPublicIP = "192.168.0.103"
	StaticPort        = "8081"

	MySQLDefaultDSN = "tik_duck:duck@tcp(localhost:9910)/tik_duck?charset=utf8mb4&parseTime=True"
	TCP             = "tcp"
	Success         = "success"
	HTTPMaxBodySize = 1024 * 1024 * 1024 // 1GB
)
