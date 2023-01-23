// 本文件放置jwt的配置信息
package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	//JWT密钥
	JWTSecret = "G808天下第一"
	//JWT定义过期时间 //目前的设定是两小时
	JWTOverTime = time.Hour * 2
)

// 目前jwt的claims中 只记录了用户名一个信息
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
