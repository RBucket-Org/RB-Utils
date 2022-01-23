package middleware

import (
	"net/http"
	"strings"

	"github.com/RBucket-Org/RB-Utils/utils/rest_errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type JwtClaim struct {
	UserID      int64
	DeviceID    int64
	IdentityKey string
	jwt.StandardClaims
}

type Validate func(token string, secret string) (*JwtClaim, rest_errors.RestError)

func ProtectedMiddleWare(extractionKey string, accessKey string, validateToken Validate, sugarLogger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get the auth token from the header of the URI
		sugarLogger.Debugf("authorization of the API starts")
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			sugarLogger.Errorf("no authorization token")
			emptyAuth := rest_errors.NewError("no authorization token", "empty_token", http.StatusForbidden)
			c.JSON(int(emptyAuth.Status()), emptyAuth)
			c.Abort()
			return
		}

		//split the extracted token
		extractedToken := strings.Split(clientToken, extractionKey)

		//check the length
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[len(extractedToken)-1])
		} else {
			sugarLogger.Errorf("invalid authorization token")
			emptyAuth := rest_errors.NewError("invalid authorization token", "empty_token", http.StatusBadRequest)
			c.JSON(int(emptyAuth.Status()), emptyAuth)
			c.Abort()
			return
		}

		//validate the token
		claims, err := validateToken(clientToken, accessKey)
		if err != nil {
			sugarLogger.Errorf(err.Message())
			c.JSON(int(err.Status()), err)
			c.Abort()
			return
		}

		//set the userID
		sugarLogger.Infof("getting the parsed ID from token for further validation")
		c.Set("user_id", claims.UserID)
		c.Set("device_id", claims.DeviceID)
		c.Set("identity_key", claims.IdentityKey)

		//execute the next handler
		c.Next()
	}
}
