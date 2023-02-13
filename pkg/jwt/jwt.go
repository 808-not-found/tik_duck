package jwt

import (
	"time"

	allerrors "github.com/808-not-found/tik_duck/pkg/allerrors"
	jwt "github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = JWTOverTime

var MySecret = []byte(JWTSecret) //nolint:gochecknoglobals

// 生成jwt //目前是只进行记录用户名.
func GenToken(username string, id int64) (string, error) {
	c := MyClaims{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "tik-duck",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 使用指定的secret签名并获得完成的编码后的字符串token
	return token.SignedString(MySecret)
}

// 解析JWT.
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i any, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, allerrors.ErrJWTParseTokenRun()
}
