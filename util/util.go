/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/8/20 17:06
 */

package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io"
	"math/rand"
	"os"
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

func Md5IO(reader io.Reader) string {
	buf := make([]byte, 1024*1024)
	h := md5.New()

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				return ""
			}
			break
		}

		if n == 0 {
			break
		}

		h.Write(buf[:n])
	}

	return hex.EncodeToString(h.Sum(nil))
}

func B64Encode(input io.Reader, output io.Writer) (err error) {

	encoder := base64.NewEncoder(base64.StdEncoding, output)
	if _, err = io.Copy(encoder, input); err != nil {
		return err
	}

	if err = encoder.Close(); err != nil {
		return err
	}

	return nil
}

func B64Decode(input io.Reader, output io.Writer) (err error) {
	decoder := base64.NewDecoder(base64.StdEncoding, input)
	if _, err = io.Copy(output, decoder); err != nil {
		return err
	}
	return nil
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

func Env(env string, dft string) string {
	v := os.Getenv(env)
	if v == "" {
		return dft
	}
	return v
}
