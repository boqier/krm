package jwtutil

import (
	"time"

	"github.com/boqier/gin-scaffold/config"
	"github.com/boqier/gin-scaffold/utils/logs"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSignKey = []byte(config.JWTSecret)

// 自定义jwt claims
type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 封装生成token函数
func GenToken(username string) (string, error) {
	claims := MyCustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.JWT_EXPIRE_TIME))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "lys",
			Subject:   "linneng",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSignKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

// 解析token
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})
	if err != nil {
		logs.Error(map[string]interface{}{"module": "jwtutil"}, "解析token失败：")
		return nil, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		logs.Warning(nil, "token无效")
		return nil, err
	}
}
