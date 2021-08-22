package rest_errors

import (
	"encoding/json"
	"errors"
	"net/http"

	// utils\mysql_errors\mysql_errors.go
	"github.com/RBucket-Org/RB-Utils/utils/mysql_errors"
)

// RestError : this hold the signature of the function
type RestError interface {
	Message() string
	Status() int64
	Code() string
}

type restError struct {
	RMessage string `json:"message"`
	RStatus  int64  `json:"status"`
	RCode    string `json:"code"`
}

func (re *restError) Status() int64 {
	return re.RStatus
}

func (re *restError) Message() string {
	return re.RMessage
}

func (re *restError) Code() string {
	return re.RCode
}

func databaseErr(err error) RestError {
	databaseErr := mysql_errors.DatabaseErrorHandler(err)
	if databaseErr.Code() != "not_mysql_error" {
		return &restError{
			RMessage: databaseErr.Message(),
			RStatus:  databaseErr.Status(),
			RCode:    databaseErr.Code(),
		}
	}

	return nil
}

//NewBadRequestError : this method implements the bad request error
func NewBadRequestError(message string, err error) RestError {
	//check mysql error
	if sqlErr := databaseErr(err); sqlErr != nil {
		return sqlErr
	}
	//final
	return &restError{
		RMessage: message,
		RStatus:  http.StatusBadRequest,
		RCode:    "bad_request",
	}
}

//NewNotFoundError : this method implements the not found error
func NewNotFoundError(message string, err error) RestError {
	//check mysql error
	if sqlErr := databaseErr(err); sqlErr != nil {
		return sqlErr
	}
	//final
	return &restError{
		RMessage: message,
		RStatus:  http.StatusNotFound,
		RCode:    "not_found",
	}
}

//NewInternalServerError : this method implements the internal server error
func NewInternalServerError(message string, err error) RestError {
	//check mysql error
	if sqlErr := databaseErr(err); sqlErr != nil {
		return sqlErr
	}
	//final
	return &restError{
		RMessage: message,
		RStatus:  http.StatusInternalServerError,
		RCode:    "internal_server_error",
	}
}

//NewUnauthorizedError : this method implements the unauthorized error
func NewUnauthorizedError(message string, err error) RestError {
	//check mysql error
	if sqlErr := databaseErr(err); sqlErr != nil {
		return sqlErr
	}
	//final
	return &restError{
		RMessage: message,
		RStatus:  http.StatusUnauthorized,
		RCode:    "unauthorized_error",
	}
}

//TokenExpired : this will return when the token get expired
func TokenExpired() RestError {
	return &restError{
		RMessage: "token expired",
		RStatus:  http.StatusExpectationFailed,
		RCode:    "token_expired",
	}
}

//NewError : this will creates the New Error
func NewError(message string, code string, status int64, err error) RestError {
	//check mysql error
	if sqlErr := databaseErr(err); sqlErr != nil {
		return sqlErr
	}
	//final
	return &restError{
		RMessage: message,
		RStatus:  status,
		RCode:    code,
	}
}

//NewRestErrorFromBytes : creates the resterror domain by taking the paramter as a slice of bytes
func NewRestErrorFromBytes(byte []byte) (RestError, error) {
	var apiErr restError
	if err := json.Unmarshal(byte, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return &apiErr, nil
}
