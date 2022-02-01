package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RBucket-Org/RB-Utils/utils/crypto_utils"
	"github.com/RBucket-Org/RB-Utils/utils/rest_errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type PublicClaim struct {
	Key string
	jwt.RegisteredClaims
}

type PublicValidate func(token string, secret string) (*PublicClaim, rest_errors.RestError)

func PublicMiddleWare(extractionKey string, restKey string, publicAuthKey string, validateToken PublicValidate, sugarLogger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get the auth token from the header of the URI
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			sugarLogger.Errorf("no authorization token")
			emptyAuth := rest_errors.NewError("no authorization token", "empty_token", http.StatusForbidden)
			c.JSON(int(emptyAuth.Status()), emptyAuth)
			c.Abort()
			return
		}

		//split the extracted token
		extractedKey := strings.Split(clientToken, extractionKey)

		if len(extractedKey) == 2 {
			clientToken = strings.TrimSpace(extractedKey[len(extractedKey)-1])
		} else {
			sugarLogger.Errorf("invalid public auth token")
			emptyAuth := rest_errors.NewError("invalid public auth token", "empty_token", http.StatusBadRequest)
			c.JSON(int(emptyAuth.Status()), emptyAuth)
			c.Abort()
			return
		}

		//validate the token
		claims, err := validateToken(clientToken, restKey)
		if err != nil {
			sugarLogger.Errorf(err.Message())
			c.JSON(int(err.Status()), err)
			c.Abort()
			return
		}

		saltByte := crypto_utils.IdentityHash.GenerateSalt(extractionKey)

		if !crypto_utils.IdentityHash.DoMatch(claims.Key, fmt.Sprintf("%s%s", extractionKey, publicAuthKey), saltByte) {
			sugarLogger.Errorf("invalid public auth key")
			emptyAuth := rest_errors.NewError("invalid public auth key", "invalid_key", http.StatusBadRequest)
			c.JSON(int(emptyAuth.Status()), emptyAuth)
			c.Abort()
			return
		}

		c.Set("middleware_type", Public)
		c.Set("user_id", int64(0))
		c.Set("device_id", "")
		c.Set("identity_key", "")
		c.Next()
	}
}
