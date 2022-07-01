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
	"reflect"
	"time"
)

func InArray[T any](value T, values []T) bool {

	vv := reflect.ValueOf(value)
	vsv := reflect.ValueOf(values)

	if vsv.Kind() != reflect.Slice {
		return false
	}

	switch vv.Kind() {
	case reflect.Interface, reflect.Func, reflect.Chan, reflect.Invalid:
		return false

	default:
		for _, i := range values {
			if reflect.DeepEqual(value, i) {
				return true
			}
		}

		return false
	}
}

func Md5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func Md5R(reader io.Reader) (hash string, err error) {
	h := md5.New()
	if _, err = io.Copy(h, reader); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func Md5RS(file io.ReadSeeker) (result string, err error) {
	if result, err = Md5R(file); err != nil {
		return "", err
	}

	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return "", err
	}
	return result, nil
}

func B64EncR(input io.Reader, output io.Writer) (err error) {

	encoder := base64.NewEncoder(base64.StdEncoding, output)
	if _, err = io.Copy(encoder, input); err != nil {
		return err
	}

	if err = encoder.Close(); err != nil {
		return err
	}

	return nil
}

func B64DecR(input io.Reader, output io.Writer) (err error) {
	decoder := base64.NewDecoder(base64.StdEncoding, input)
	if _, err = io.Copy(output, decoder); err != nil {
		return err
	}
	return nil
}

func B64EncRS(input io.ReadSeeker, output io.Writer) (err error) {
	if err = B64EncR(input, output); err != nil {
		return
	}

	_, err = input.Seek(0, io.SeekStart)
	return
}

func B64DecRS(input io.ReadSeeker, output io.Writer) (err error) {
	if err = B64DecR(input, output); err != nil {
		return
	}
	_, err = input.Seek(0, io.SeekStart)
	return
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

// func Md5R(reader io.Reader) string {
// 	buf := make([]byte, 1024*1024)
// 	h := md5.New()

// for {
// 	n, err := reader.Read(buf)
// 	if err != nil {
// 		if err != io.EOF {
// 			return ""
// 		}
// 		break
// 	}
//
// 	if n == 0 {
// 		break
// 	}
//
// 	h.Write(buf[:n])
// }

// 	return hex.EncodeToString(h.Sum(nil))
// }
