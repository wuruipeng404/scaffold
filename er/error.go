/*
* @Author: Rumple
* @Email: ruipeng.wu@cyclone-robotics.com
* @DateTime: 2022/1/28 11:22
 */

package er

type IError interface {
	Code() int
	Message() string
}

type _Error struct {
	code    int
	message string
}

func (e _Error) Code() int {
	return e.code
}

func (e _Error) Message() string {
	return e.message
}

func New(code int, message string) IError {
	return &_Error{
		code:    code,
		message: message,
	}
}
