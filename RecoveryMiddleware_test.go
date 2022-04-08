package gore

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMiddleware(t *testing.T) {
	r := gin.Default()

	r.Use(RecoveryHandlerMiddleware())

	r.GET("test", func(ctx *gin.Context) {
		panic("Some Error")
	})

	r.Run()
}
