// 本文件放置jwt的配置信息
package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	JWTSecret = "password_tik_duck"

	JWTOverTime = time.Hour * 2
)

// 目前jwt的claims中 只记录了用户名一个信息.
type MyClaims struct {
	ID       int64
	Username string `json:"username"`
	jwt.StandardClaims
}
