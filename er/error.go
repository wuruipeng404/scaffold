/*
* @Author: Rumple
* @Email: ruipeng.wu@cyclone-robotics.com
* @DateTime: 2022/1/28 11:22
 */

package er

import (
	"fmt"
)

type IError interface {
	Code() int
	Message() string
	Error() string
	AddArgs(args ...any)
}

type _Error struct {
	error
	code        int
	message     string
	messageArgs []any
}

func (e _Error) Code() int {
	return e.code
}

func (e _Error) Message() string {
	if len(e.messageArgs) == 0 {
		return e.message
	} else {
		return fmt.Sprintf(e.message, e.messageArgs...)
	}
}

func (e _Error) Error() string {
	return fmt.Sprintf(`[%d] %s`, e.Code(), e.Message())
}

func (e _Error) AddArgs(args ...any) {
	e.messageArgs = append(e.messageArgs, args...)
}

func New(code int, message string, args ...any) IError {
	return &_Error{
		code:        code,
		message:     message,
		messageArgs: args,
	}
}
