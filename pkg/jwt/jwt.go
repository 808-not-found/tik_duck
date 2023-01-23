package jwt

import (
	"errors"
	"time"

	setting "github.com/808-not-found/tik_duck/pkg/jwt/setting"
	jwt "github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = setting.JWTOverTime

var MySecret = []byte(setting.JWTSecret)

// 生成jwt //目前是只进行记录用户名
func GenToken(username string) (string, error) {
	c := setting.MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "tik-duck",
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	//使用指定的secret签名并获得完成的编码后的字符串token
	return token.SignedString(MySecret)
}

// 解析JWT
func ParseToken(tokenString string) (*setting.MyClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &setting.MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
