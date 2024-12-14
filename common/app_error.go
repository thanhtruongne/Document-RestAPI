package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	EntityName = "Test"
)

var (
	RecordNotFound     = errors.New("record not found")
	RecordCannotCreate = errors.New("record cannot create")
)

type AppError struct {
	StatusCode int    `json:"status_code"` //  mã lỗi 400,404,500/.....
	RootErr    error  `json:"-"`           // lỗi gốc
	Message    string `json:"message"`     //  lỗi thông báo client
	Log        string `json:"log"`         // lỗi data lấy từ rooterro
	Key        string `json:"key"`         //  custom lỗi đa ngôn ngữ
}

func NewFullErrorResponse(statusCode int, root error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewErrorReponse(root error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewErrorCustomResponse(root error, message, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    message,
		Key:        key,
	}
}

func NewErrorAutheorizedReponse(root error, message, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    message,
		Key:        key,
	}
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError() // trả về chính nó loop theo bọc theo trỏ
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

//custom các lỗi

func ErrorDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Error on database like", err.Error(), "DB_ERROR")
}
func ErrorInvalidRequest(err error) *AppError {
	return NewErrorReponse(err, "Invalid Request ", err.Error(), "INVALID_REQUEST")
}
func ErrorInternalServerError(err error) *AppError {
	return NewErrorReponse(err, "Internal Server Error ", err.Error(), "INTERNAL_SERVER_ERROR")
}

func ErrCanNotListEntity(entity string, err error) *AppError {
	return NewErrorCustomResponse(
		err,
		fmt.Sprintf("Cannot List data %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCanNotDeleteEntity(entity string, err error) *AppError {
	return NewErrorCustomResponse(
		err,
		fmt.Sprintf("Cannot List data %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCanNotCreateEntity(entity string, err error) *AppError {
	return NewErrorCustomResponse(
		err,
		fmt.Sprintf("Cannot List data %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}
