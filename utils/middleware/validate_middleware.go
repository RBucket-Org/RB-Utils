package middleware

import (
	"fmt"

	"github.com/RBucket-Org/RB-Utils/utils/rest_errors"
	"github.com/gin-gonic/gin"
)

type MiddlewareValue struct {
	MiddlewareType MiddlewareType
	UserID         int64
	DeviceID       string
	IdentityKey    string
}

func (mv *MiddlewareValue) GetMiddlewareValue(ctx *gin.Context) rest_errors.RestError {
	//if the request is from public api then bypass the validation
	middleWareType, check := ctx.Get("middleware_type")
	if !check {
		err := rest_errors.NewBadRequestError("token is invalid")
		return err
	}
	//get the user_id from the middleware
	userID, check := ctx.Get("user_id")
	if !check {
		err := rest_errors.NewBadRequestError("token is invalid")
		return err
	}
	//get the device_id from the middleware
	deviceID, check := ctx.Get("device_id")
	if !check {
		err := rest_errors.NewBadRequestError("token is invalid")
		return err
	}
	//get the identity_key from the middleware
	identityKey, check := ctx.Get("identity_key")
	if !check {
		err := rest_errors.NewBadRequestError("token is invalid")
		return err
	}
	fmt.Println("userID, deviceID, identityKey", userID, deviceID, identityKey)

	mv.UserID = userID.(int64)
	mv.DeviceID = deviceID.(string)
	mv.IdentityKey = identityKey.(string)
	mv.MiddlewareType = middleWareType.(MiddlewareType)

	return nil
}