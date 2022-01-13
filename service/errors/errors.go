// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Original source: github.com/micro/go-micro/v3/errors/errors.go

// Package errors provides a way to return detailed information
// for an RPC request error. The error is normally JSON encoded.
package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	RequestId string `json:"request_id,omitempty"`
	Code      int32  `json:"code,omitempty"`
	Detail    string `json:"detail,omitempty"`
	Status    string `json:"status,omitempty"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// New generates a custom error.
func New(detail string, code int32) error {
	return &Error{
		Code:   code,
		Detail: detail,
		Status: http.StatusText(int(code)),
	}
}

// FromError try to convert go error to *Error
func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if verr, ok := err.(*Error); ok && verr != nil {
		return verr
	}

	return Parse(err.Error())
}

// Parse tries to parse a JSON string into an error. If that
// fails, it will set the given string as the error detail.
func Parse(err string, id ...string) *Error {
	e := new(Error)
	if e1 := json.Unmarshal([]byte(err), e); e1 != nil {
		e.Detail = err
	}
	if len(id) > 0 && len(e.RequestId) == 0 {
		e.RequestId = id[0]
	}
	return e
}

// BadRequest generates a 400 error.
func BadRequest(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusBadRequest,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(400),
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusUnauthorized,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusUnauthorized),
	}
}

// Forbidden generates a 403 error.
func Forbidden(format string, a ...interface{}) error {
	return &Error{
		Code:   403,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusForbidden),
	}
}

// NotFound generates a 404 error.
func NotFound(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusNotFound,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusNotFound),
	}
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusMethodNotAllowed,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusMethodNotAllowed),
	}
}

// Timeout generates a 408 error.
func Timeout(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusRequestTimeout,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusRequestTimeout),
	}
}

// Conflict generates a 409 error.
func Conflict(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusConflict,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusConflict),
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusInternalServerError,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusInternalServerError),
	}
}

// NotImplemented generates a 501 error
func NotImplemented(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusNotImplemented,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusNotImplemented),
	}
}

// BadGateway generates a 502 error
func BadGateway(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusBadGateway,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusBadGateway),
	}
}

// ServiceUnavailable generates a 503 error
func ServiceUnavailable(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusServiceUnavailable,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusServiceUnavailable),
	}
}

// GatewayTimeout generates a 504 error
func GatewayTimeout(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusGatewayTimeout,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusGatewayTimeout),
	}
}

// Equal tries to compare errors
func Equal(err1 error, err2 error) bool {
	verr1, ok1 := err1.(*Error)
	verr2, ok2 := err2.(*Error)

	if ok1 != ok2 {
		return false
	}

	if !ok1 {
		return err1 == err2
	}

	if verr1.Code != verr2.Code {
		return false
	}

	return true
}
