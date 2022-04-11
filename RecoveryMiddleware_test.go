package gore

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMiddleware(t *testing.T) {
	r := gin.Default()

	r.Use(RecoveryHandlerMiddleware())

	r.GET("test", func(ctx *gin.Context) {
		panic("Some Error")
	})

	r.GET("expected", func(ctx *gin.Context) {
		panic(NewException(fmt.Errorf("Expected wrapped error"), "Error Occured", "An expected error occured", "", "", 400, 10000))
	})

	r.Run()
}
