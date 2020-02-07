package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey = []byte("zhengyansheng")
)

/**
 * 生成 token
 */
func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      int64(time.Now().Add(time.Minute * 20).Unix()), // 可以添加过期时间
	})

	//对应的字符串请自行生成，最后足够使用加密后的字符串
	return token.SignedString(secretKey)
}

/**
 * 解析 token
 */
func ParseToken(tokenSrt string) (string, error) {
	token, err := jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "jwt.Parse error.", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if u, ok := claims["username"].(string); ok {
			return u, nil
		} else {
			return "", errors.New("断言错误")
		}
	}
	return "", errors.New("Token.Valid error")
}
