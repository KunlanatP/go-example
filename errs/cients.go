package errs

import "net/http"

func Invalid(message ...string) error {
	return newError(http.StatusBadRequest, message)
}
