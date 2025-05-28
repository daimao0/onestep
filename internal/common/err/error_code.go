package err

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
	Success        ErrorCode = ErrorCode{Code: 200, Msg: "success", Err: nil}
	Fail           ErrorCode = ErrorCode{Code: 500, Msg: "fail", Err: nil}
	ValidateFailed ErrorCode = ErrorCode{Code: 400, Msg: "validate fail", Err: nil}
	Unauthorized   ErrorCode = ErrorCode{Code: 401, Msg: "unauthorized", Err: nil}
)

func (e *ErrorCode) Error() string {
	bytes, err := json.MarshalIndent(e, "", "  ") // 使用两个空格作为缩进
	if err != nil {
		// 处理json.MarshalIndent可能发生的错误
		return fmt.Sprintf("{\"error\": \"%v\"}", err)
	}
	return string(bytes)
}
