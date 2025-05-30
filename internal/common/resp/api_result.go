package resp

import "onestep/internal/common/error_code"

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
		Code:    error_code.Fail.Code,
		Message: message,
	}
}

// Success return success resp
func Success(data interface{}) Resp {
	return Resp{
		Code: error_code.Success.Code,
		Data: data,
	}
}

// Unauthorized return unauthorized resp
func Unauthorized() Resp {
	return Resp{
		Code:    error_code.Unauthorized.Code,
		Message: error_code.Unauthorized.Msg,
	}
}

// InvalidParam return invalidParam resp
func InvalidParam(msg string) Resp {
	return Resp{
		Code:    error_code.ValidateFailed.Code,
		Message: error_code.ValidateFailed.Msg + "; " + msg,
	}
}

// SystemError return system error resp
func SystemError(errorCode *error_code.ErrorCode) Resp {
	return Resp{
		Code:    errorCode.Code,
		Message: errorCode.Msg,
		Data:    nil,
	}
}
