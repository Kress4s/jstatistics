package handlers

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type BaseHandler struct {
	UserID   uint
	UserName string
}

func (bh *BaseHandler) BeginRequest(ctx iris.Context) {
	token := ctx.Values().Get("jwt").(*jwt.Token)
	userInfo := token.Claims.(jwt.MapClaims)
	id := userInfo["user_id"].(float64)
	bh.UserID = uint(id)
	bh.UserName = userInfo["user_name"].(string)
}

func (bh *BaseHandler) EndRequest(ctx iris.Context) {}
