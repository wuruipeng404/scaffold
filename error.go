/*
* @Author: Rumple
* @Email: ruipeng.wu@cyclone-robotics.com
* @DateTime: 2022/1/28 11:22
 */

package scaffold

type IError interface {
	Code() int
	Message() string
}

type SError struct {
	code    int
	message string
}

func (e SError) Code() int {
	return e.code
}

func (e SError) Message() string {
	return e.message
}

func NewError(code int, message string) IError {
	return &SError{
		code:    code,
		message: message,
	}
}
