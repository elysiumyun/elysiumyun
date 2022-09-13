package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func SetMiddleware(engine *gin.Engine) {
	engine.Use(ginLogger())
}

func ginLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s\"%s [\" %s\"]\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage)
	})
}
