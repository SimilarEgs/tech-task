package httperrors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

type APIerr interface {
	Status() int
	Error() string
	Causes() interface{}
	Render(w http.ResponseWriter, r *http.Request) error
}

var (
	NotFound            = errors.New("user_not_found")
	BadRequest          = errors.New("bad_request")
	InternalServerError = errors.New("internal_server_error")
)

type APIerror struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCauses interface{} `json:"-"`
}

func (e APIerror) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.ErrStatus)
	return nil
}

func (e APIerror) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e APIerror) Status() int {
	return e.ErrStatus
}

func (e APIerror) Causes() interface{} {
	return e.ErrCauses
}

func NewAPIerror(status int, err string, causes interface{}) APIerr {
	return APIerror{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

func NewBadRequestError(causes interface{}) APIerr {
	return APIerror{
		ErrStatus: http.StatusBadRequest,
		ErrError:  BadRequest.Error(),
		ErrCauses: causes,
	}
}

func NewNotFoundError(causes interface{}) APIerr {
	return APIerror{
		ErrStatus: http.StatusNotFound,
		ErrError:  NotFound.Error(),
		ErrCauses: causes,
	}
}

func NewInternalServerError(causes interface{}) APIerr {
	result := APIerror{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  InternalServerError.Error(),
		ErrCauses: causes,
	}
	return result
}

func ParseErrors(err error) APIerr {
	switch {
	case errors.Is(err, NotFound):
		return NewAPIerror(http.StatusNotFound, NotFound.Error(), err)
	case strings.Contains(err.Error(), "Unmarshal"):
		return NewAPIerror(http.StatusBadRequest, BadRequest.Error(), err)
	case strings.Contains(err.Error(), "empty"):
		return NewAPIerror(http.StatusBadRequest, BadRequest.Error(), err)
	default:
		return NewInternalServerError(err)
	}
}

func ErrorResponse(err error) interface{} {
	return ParseErrors(err)
}
