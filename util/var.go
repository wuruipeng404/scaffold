/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/8/23 14:45
 */

package util

const (
	TimeFormatString   = "2006-01-02 15:04:05.000"
	TimeFormatStandard = "2006-01-02 15:04:05"
	RandomStrSeed      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	RandomStrLowerSeed = "abcdefghijklmnopqrstuvwxyz0123456789"
)

var (
	letters = []rune(RandomStrSeed)
)

type SwagBase struct {
	Code int    `json:"code" example:"0"`      // 响应码 非0即为失败
	Msg  string `json:"msg" example:"success"` // msg
}
