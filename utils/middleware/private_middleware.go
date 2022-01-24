package middleware

import "github.com/gin-gonic/gin"

type HandleMiddleWare func(ctx *gin.Context)

func PrivateMiddleware(passcode string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//authorize the private API

		c.Set("middleware_type", Private)
		c.Next()
	}
}
