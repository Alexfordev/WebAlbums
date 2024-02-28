package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
)

var (
	ParamsRequestTooManyError = map[string]interface{}{"code": 429, "msg": "requests too many"}
)

func LimitHandler(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {

		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.JSON(http.StatusOK, ParamsRequestTooManyError)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
