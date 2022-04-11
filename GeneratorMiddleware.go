package gore

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//for every transaction(request) create a new transactionID in the header
//if correlationID empty, create a new corrID
func GeneratorHandlerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			correlationId := ctx.Request.Header.Get("CorrelationId")
			ctx.Request.Header.Set("TransactionId", uuid.New().String())

			if correlationId == "" {
				ctx.Request.Header.Set("CorrelationId", uuid.New().String())
			}

			ctx.JSON(200, "Transaction, correlation ID checked")

		}()
		ctx.Next()
	}
}
