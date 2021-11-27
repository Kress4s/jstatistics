package middlewares

import (
	"js_statistics/app/service"
	"js_statistics/app/vo"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12/context"
)

func RecordSystemLog(funcName, param, content string) context.Handler {
	return func(ctx *context.Context) {
		token := ctx.Values().Get("jwt").(*jwt.Token)
		userInfo := token.Claims.(jwt.MapClaims)
		userName := userInfo["user_name"].(string)
		ip := ctx.RemoteAddr()
		address := ctx.Path()
		switch funcName {
		case "Update":
			id := ctx.Params().GetString(param)
			content = content + param + ": " + id
		case "Delete":
			id := ctx.Params().GetString(param)
			content = content + param + ": " + id
		case "MultiDelete":
			ids := ctx.URLParam(param)
			content = content + param + ": " + ids
		}
		if err := service.GetSyslogService().Create(&vo.SystemLogReq{
			UserName:    userName,
			IP:          ip,
			Address:     address,
			Description: content,
		}); err != nil {
			ctx.Application().Logger().Error("failed to record system log: ", err)
		}
		ctx.Next()
	}
}
