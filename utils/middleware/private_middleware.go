package middleware

import "github.com/gin-gonic/gin"

type HandleMiddleWare func(ctx *gin.Context)

func PrivateMiddleware(passcode string) HandleMiddleWare {
	return func(c *gin.Context) {
		//authorize the private API
	}
}
