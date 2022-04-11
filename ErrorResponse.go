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
	GatewayError
	MethodNotAllowedError
	ServerOverburdenedError
	GatewayTimeoutError
	InternalServalError
	EntityGoneError
	EntityConflictError
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

func GatewayException(e error, transactionId, correlationId string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:    http.StatusBadGateway,
		ErrorCode:     GatewayError,
		ErrorMessage:  "GatewayError",
		TransactionId: transactionId,
		CorrelationId: correlationId,
		ErrorDetails:  e.Error(),
		err:           e,
	}
}

func MethodNotAllowedException(e error, transactionId, correlationId string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:    http.StatusMethodNotAllowed,
		ErrorCode:     MethodNotAllowedError,
		ErrorMessage:  "MethodNotAllowedError",
		TransactionId: transactionId,
		CorrelationId: correlationId,
		ErrorDetails:  e.Error(),
		err:           e,
	}
}

func ServerOverburdenedException(e error, transactionId, correlationId string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:    http.StatusServiceUnavailable,
		ErrorCode:     ServerOverburdenedError,
		ErrorMessage:  "ServerOverburdenedError", //this status usually means server is overburdened with requests or under maintenance
		TransactionId: transactionId,
		CorrelationId: correlationId,
		ErrorDetails:  e.Error(),
		err:           e,
	}
}

func GatewayTimeoutException(e error, transactionId, correlationId string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:    http.StatusGatewayTimeout,
		ErrorCode:     GatewayTimeoutError,
		ErrorMessage:  "GatewayTimeoutError",
		TransactionId: transactionId,
		CorrelationId: correlationId,
		ErrorDetails:  e.Error(),
		err:           e,
	}
}

func InternalServerException(e error, transactionId, correlationId string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:    http.StatusInternalServerError,
		ErrorCode:     InternalServalError, //check server logs
		ErrorMessage:  "InternalServalError",
		TransactionId: transactionId,
		CorrelationId: correlationId,
		ErrorDetails:  e.Error(),
		err:           e,
	}
}

func EntityGoneException(e error, transactionId, correlationId string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:    http.StatusGone,
		ErrorCode:     EntityGoneError, //resource no longer available
		ErrorMessage:  "EntityGoneError",
		TransactionId: transactionId,
		CorrelationId: correlationId,
		ErrorDetails:  e.Error(),
		err:           e,
	}
}

func EntityConflictException(e error, transactionId, correlationId string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:    http.StatusConflict,
		ErrorCode:     EntityConflictError,
		ErrorMessage:  "EntityConflictError", //a request conflicts with current resource state, updates, versions
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
