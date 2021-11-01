/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/10/11 14:40
 */

package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func JwtCreateToken(secret, sub string, expire time.Duration) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expire).Unix(), // 过期时间
		Id:        uuid.New().String(),           // 唯一ID
		IssuedAt:  time.Now().Unix(),             // 发行时间
		Subject:   sub,                           // 用户信息
	}).SignedString([]byte(secret))
}

func JwtParseToken(secret, token string) (sub string, err error) {
	var (
		claim *jwt.Token
	)

	if claim, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}); err != nil {
		return
	}

	sub = claim.Claims.(jwt.MapClaims)["sub"].(string)
	return
}
