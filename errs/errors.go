package errs

import "net/http"

func newError(code int, message []string) *Error {
	var msg string
	if len(message) > 0 {
		msg = message[0]
	} else {
		msg = http.StatusText(code)
	}
	return &Error{
		Code:    code,
		Message: msg,
	}
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return err.Message
}

var ErrIDRequired = Invalid("ID is required")
