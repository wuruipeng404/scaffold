/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/8/23 14:45
 */

package util

const (
	TimeFormatStandard = "2006-01-02 15:04:05"
	TimeFormatMill     = TimeFormatStandard + ".000"
	RandomStrNumber    = "0123456789"
	RandomStrLowerSeed = "abcdefghijklmnopqrstuvwxyz" + RandomStrNumber
	RandomStrSeed      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + RandomStrLowerSeed
)

var (
	letters = []rune(RandomStrSeed)
)
