/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/8/23 14:45
 */

package util

const (
	TimeFormatStandard = "2006-01-02 15:04:05"
	TimeFormatMill     = "2006-01-02 15:04:05.000"
	RandomStrSeed      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	RandomStrLowerSeed = "abcdefghijklmnopqrstuvwxyz0123456789"
)

var (
	letters = []rune(RandomStrSeed)
)
