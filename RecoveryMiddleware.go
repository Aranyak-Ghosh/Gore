package gore

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

func RecoveryHandlerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var ne *ErrorResponse
				var ok bool = false
				var temp ErrorResponse
				if reflect.ValueOf(r).Kind() == reflect.Pointer && r != nil {
					ne, ok = r.(*ErrorResponse)
				} else {
					temp, ok = r.(ErrorResponse)
					ne = &temp
				}

				transactionId := ctx.Request.Header.Get("TransactionId")
				correlationId := ctx.Request.Header.Get("CorrelationId")

				if ne.TransactionId == "" {
					ne.TransactionId = transactionId
				}

				if ne.CorrelationId == "" {
					ne.CorrelationId = correlationId
				}

				if ok {
					if ne.StatusCode == 0 {
						ne.StatusCode = 500
					}
				} else {
					ne = &ErrorResponse{
						StatusCode:    500,
						ErrorCode:     9999,
						ErrorMessage:  "UnhandledError",
						ErrorDetails:  fmt.Sprintf("%v", r),
						TransactionId: transactionId,
						CorrelationId: correlationId,
					}
				}

				ctx.JSON(ne.StatusCode, ne)
			}
		}()
		ctx.Next()
	}
}
