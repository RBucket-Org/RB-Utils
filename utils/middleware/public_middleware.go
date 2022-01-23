package middleware

import (
	"crypto/sha512"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/RBucket-Org/RB-Utils/utils/rest_errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PublicMiddleWare(extractionKey string, publicAuthKey string, sugarLogger *zap.SugaredLogger) gin.HandlerFunc {
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

		hash := sha512.New()
		hash.Write([]byte(clientToken))
		val := hex.EncodeToString(hash.Sum(nil))
		//get the public auth key
		generateFromAuthKey := val

		if generateFromAuthKey != publicAuthKey {
			sugarLogger.Errorf("invalid public auth key")
			emptyAuth := rest_errors.NewError("invalid public auth key", "invalid_key", http.StatusBadRequest)
			c.JSON(int(emptyAuth.Status()), emptyAuth)
			c.Abort()
			return
		}

		c.Next()
	}
}
