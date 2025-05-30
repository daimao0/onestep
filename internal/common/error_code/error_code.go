package error_code

import (
	"encoding/json"
	"fmt"
)

// ErrorCode is error code common struct
type ErrorCode struct {
	Code int
	Msg  string
	Err  error
}

var (
	// Success common success code
	Success = ErrorCode{Code: 200, Msg: "success", Err: nil}
	// Fail common fail code
	Fail = ErrorCode{Code: 500, Msg: "fail", Err: nil}
	// ValidateFailed common validate fail code
	ValidateFailed = ErrorCode{Code: 400, Msg: "validate fail", Err: nil}
	// Unauthorized common unauthorized code
	Unauthorized = ErrorCode{Code: 401, Msg: "unauthorized", Err: nil}

	//---------------------------------------------------
	//  code source error custom  code 1001-1999
	//-----------------------------------------------------

	// CodeSourceInitFailed code source error code 1001-1999
	CodeSourceInitFailed = ErrorCode{Code: 10001, Msg: "code source init failed", Err: nil}
)

// NewErrorCode wraps an error with a error code
func NewErrorCode(errorCode ErrorCode, err error) *ErrorCode {
	return &ErrorCode{Code: errorCode.Code, Msg: errorCode.Msg, Err: err}
}

func (e *ErrorCode) Error() string {
	bytes, err := json.MarshalIndent(e, "", "  ") // 使用两个空格作为缩进
	if err != nil {
		// 处理json.MarshalIndent可能发生的错误
		return fmt.Sprintf("{\"error\": \"%v\"}", err)
	}
	return string(bytes)
}
