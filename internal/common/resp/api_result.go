package resp

import "onestep/internal/common/err"

// Resp struct
type Resp struct {
	// Code api status
	Code int `json:"code"`
	// Message api message
	Message string `json:"message"`
	// Data api data
	Data interface{} `json:"data"`
}

// Fail return fail resp
func Fail(message string) Resp {
	return Resp{
		Code:    err.Fail.Code,
		Message: message,
	}
}

// Success return success resp
func Success(data interface{}) Resp {
	return Resp{
		Code: err.Success.Code,
		Data: data,
	}
}

// Unauthorized return unauthorized resp
func Unauthorized() Resp {
	return Resp{
		Code:    err.Unauthorized.Code,
		Message: err.Unauthorized.Msg,
	}
}

// InvalidParam return invalidParam resp
func InvalidParam(msg string) Resp {
	return Resp{
		Code:    err.ValidateFailed.Code,
		Message: err.ValidateFailed.Msg + "; " + msg,
	}
}

// SystemError return system error resp
func SystemError(errorCode *err.ErrorCode) Resp {
	return Resp{
		Code:    errorCode.Code,
		Message: errorCode.Msg,
		Data:    nil,
	}
}
