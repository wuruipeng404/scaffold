/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/8/20 17:06
 */

package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"mime/multipart"
	"time"
)

type Element interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | string |
	~float32 | ~float64 | bool
}

func InArray[T Element](value T, values []T) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func Md5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// FileMd5 todo: large file hash
func FileMd5(file multipart.File) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	// reset reader seek
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func RandomStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// UTCNow for database time field
func UTCNow() time.Time {
	return time.Now().UTC().Truncate(time.Millisecond)
}

func IF(env string, def string) string {
	if env == "" {
		return def
	}
	return env
}

// // IFNum from env get int value // todo: Generic
// func IFNum(env string, defaultValue int) int {
// 	if env == "" {
// 		return defaultValue
// 	}
//
// 	if value, err := strconv.Atoi(env); err != nil {
// 		return defaultValue
// 	} else {
// 		return value
// 	}
// }
