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
		err := rest_errors.NewBadRequestError("token is invalid because the middleware type is not valid")
		return err
	}
	//get the user_id from the middleware
	userID, check := ctx.Get("user_id")
	if !check {
		err := rest_errors.NewBadRequestError("token is invalid because the user id is not valid")
		return err
	}
	//get the device_id from the middleware
	deviceID, check := ctx.Get("device_id")
	if !check {
		err := rest_errors.NewBadRequestError("token is invalid because the device id is not valid")
		return err
	}
	//get the identity_key from the middleware
	identityKey, check := ctx.Get("identity_key")
	if !check {
		err := rest_errors.NewBadRequestError("token is invalid because the identity key is not valid")
		return err
	}
	fmt.Printf("userID:%s\n, deviceID:%s\n, identityKey:%s\n, type:%s\n", userID, deviceID, identityKey, middleWareType)

	mv.UserID = userID.(int64)
	mv.DeviceID = deviceID.(string)
	mv.IdentityKey = identityKey.(string)
	mv.MiddlewareType = middleWareType.(MiddlewareType)

	return nil
}
