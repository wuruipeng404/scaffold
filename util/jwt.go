/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/10/11 14:40
 */

package util

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

func JwtCreateToken(secret, sub string, expire time.Duration) (string, error) {

	return jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)), // 过期时间
		ID:        uuid.New().String(),                        // 唯一ID
		IssuedAt:  jwt.NewNumericDate(time.Now()),             // 发行时间
		Subject:   sub,                                        // 用户信息
	}).SignedString([]byte(secret))

	// return jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
	// 	ExpiresAt: time.Now().Add(expire).Unix(), // 过期时间
	// 	Id:        uuid.New().String(),           // 唯一ID
	// 	IssuedAt:  time.Now().Unix(),             // 发行时间
	// 	Subject:   sub,                           // 用户信息
	// }).SignedString([]byte(secret))
}

func JwtParseToken(secret, token string) (sub string, err error) {
	var (
		claim *jwt.Token
	)

	if claim, err = jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	}); err != nil {
		return
	}

	sub = claim.Claims.(jwt.MapClaims)["sub"].(string)
	return
}

func JwtParseInfo(token string) (result jwt.MapClaims, err error) {

	claim, err := jwt.Parse(token, nil)
	if claim == nil {
		return nil, err
	}

	return claim.Claims.(jwt.MapClaims), nil
}
