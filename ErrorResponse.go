package gore

import (
	"net/http"
)

type ErrorResponse struct {
	StatusCode    int    `json:"-"`
	ErrorCode     int    `json:"errorCode,omitEmpty"`
	ErrorMessage  string `json:"errorMessage,omitEmpty"`
	ErrorDetails  string `json:"errorDetails,omitEmpty"`
	TransactionId string `json:"transactionId,omitEmpty"`
	CorrelationId string `json:"correlationId,omitEmpty"`
	err           error  `json:"-"`
}

type ErrorCode = int

const (
	EntityNotFoundError ErrorCode = 1000 + iota
	EntityValidationError
	EntityUnauthorizedError
	TransactionForbiddenError
)

func (err *ErrorResponse) Error() string {
	if err.err != nil {
		return err.err.Error()
	} else {
		return err.ErrorDetails
	}
}

func (err *ErrorResponse) Unwrap() error {
	return err.err
}

func EntityValidationException(e error, transactionId, correlationId string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:    http.StatusBadRequest,
		ErrorCode:     EntityValidationError,
		ErrorMessage:  "EntityValidationError",
		TransactionId: transactionId,
		CorrelationId: correlationId,
		ErrorDetails:  e.Error(),
		err:           e,
	}
}

func EntityNotFoundException(e error, transactionId, correlationId string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:    http.StatusNotFound,
		ErrorCode:     EntityNotFoundError,
		ErrorMessage:  "EntityNotFoundError",
		TransactionId: transactionId,
		CorrelationId: correlationId,
		ErrorDetails:  e.Error(),
		err:           e,
	}
}

func EntityUnauthorizedException(e error, transactionId, correlationId string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:    http.StatusUnauthorized,
		ErrorCode:     EntityUnauthorizedError,
		ErrorMessage:  "EntityUnauthorizedError",
		TransactionId: transactionId,
		CorrelationId: correlationId,
		ErrorDetails:  e.Error(),
		err:           e,
	}
}

func TransactionForbiddenException(e error, transactionId, correlationId string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:    http.StatusForbidden,
		ErrorCode:     TransactionForbiddenError,
		ErrorMessage:  "TransactionForbiddenError",
		TransactionId: transactionId,
		CorrelationId: correlationId,
		ErrorDetails:  e.Error(),
		err:           e,
	}
}

func NewException(e error, errorMessage, errorDetails, transactionId, correlationId string, statusCode, errorCoder int) *ErrorResponse {
	return &ErrorResponse{
		err:           e,
		TransactionId: transactionId,
		CorrelationId: correlationId,
		StatusCode:    statusCode,
		ErrorCode:     errorCoder,
		ErrorMessage:  errorMessage,
		ErrorDetails:  errorDetails,
	}
}
