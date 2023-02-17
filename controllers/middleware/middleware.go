package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Middleware99() gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.GetHeader("HeaderAuthorization")

		if header != "mysecretpassword" {
			c.AbortWithError(401, fmt.Errorf("%s", "this is a private feature, please provide the token authorization"))
		}

		c.Next()
	}
}
