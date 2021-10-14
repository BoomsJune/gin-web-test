package util

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"example.com/web-test/config"
	"github.com/golang-jwt/jwt"
)

const (
	headerName = "Authorization" // 需要包含的JWT请求头名称
	headerType = "Bearer"        // JWT请求头类型，以此开头
	expires    = time.Hour * 1   // JWT有效时长
	issuer     = "my-app"        // 签发者
)

// 转换下secret格式，原因见源码https://github.com/dgrijalva/jwt-go/blob/v3.2.0/hmac.go#L82
func getHS256Secret() []byte {
	return []byte(config.Cfg.JWT.Secret)
}

// 生成jwt
func JwtGet(identity uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expires).Unix(),
		Subject:   fmt.Sprintf("%d", identity),
		Issuer:    issuer,
	})
	return token.SignedString(getHS256Secret())
}

// 解析jwt
func JwtParse(header http.Header) (*jwt.StandardClaims, error) {
	parts := strings.SplitN(header.Get(headerName), " ", 2)
	if len(parts) != 2 || parts[0] != headerType {
		return nil, errors.New("invalid auth token")
	}

	token, err := jwt.ParseWithClaims(parts[1], &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return getHS256Secret(), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims, nil
}
