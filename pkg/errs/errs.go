package errs

import "net/http"

type Err struct {
	StatusCode int    `json:"status"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

func BadRequest(message string, err error) *Err {
	return &Err{
		StatusCode: http.StatusBadRequest,
		Error:      err.Error(),
		Message:    message,
	}
}

func Unauthorized(message string, err error) *Err {
	return &Err{
		StatusCode: http.StatusUnauthorized,
		Error:      err.Error(),
		Message:    message,
	}
}

func Forbidden(message string, err error) *Err {
	return &Err{
		StatusCode: http.StatusForbidden,
		Error:      err.Error(),
		Message:    message,
	}
}

func NotFound(message string, err error) *Err {
	return &Err{
		StatusCode: http.StatusNotFound,
		Error:      err.Error(),
		Message:    message,
	}
}

func Conflict(message string, err error) *Err {
	return &Err{
		StatusCode: http.StatusConflict,
		Error:      err.Error(),
		Message:    message,
	}
}

func InternalServerError(message string, err error) *Err {
	return &Err{
		StatusCode: http.StatusInternalServerError,
		Error:      err.Error(),
		Message:    message,
	}
}

func BadGateway(message string, err error) *Err {
	return &Err{
		StatusCode: http.StatusBadGateway,
		Error:      err.Error(),
		Message:    message,
	}
}

func ServiceUnavailable(message string, err error) *Err {
	return &Err{
		StatusCode: http.StatusServiceUnavailable,
		Error:      err.Error(),
		Message:    message,
	}
}

func GatewayTimeout(message string, err error) *Err {
	return &Err{
		StatusCode: http.StatusServiceUnavailable,
		Error:      err.Error(),
		Message:    message,
	}
}
