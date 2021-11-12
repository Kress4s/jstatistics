package routes

import (
	"js_statistics/app/handlers/auth"
	v1 "js_statistics/app/handlers/v1"
	"js_statistics/app/middlewares"
	"js_statistics/config"
	_ "js_statistics/docs"

	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterRoutes(app *iris.Application) {
	cfg := config.GetConfig()
	if cfg.DebugModel {
		app.Get("/swagger/{any:path}", swagger.WrapHandler(swaggerFiles.Handler))
	}
	app.Get("/liveness", func(ctx iris.Context) {
		ctx.ResponseWriter().WriteHeader(iris.StatusOK)
	})

	authApp := app.Party("/auth")
	mvc.New(authApp).Handle(auth.NewLoginHandler())

	// 权限管理
	party := app.Party("/api/v1")
	party.Use(middlewares.Auth().Serve)
	permissionParty := party.Party("/permission")
	permissionApp := mvc.New(permissionParty)
	permissionApp.Handle(v1.NewUserHandler())
	permissionApp.Handle(v1.NewPermissionHandler())
	permissionApp.Handle(v1.NewRoleHandler())

	// 应用管理
	applicationParty := party.Party("/application")
	applicationApp := mvc.New(applicationParty)
	applicationApp.Handle(v1.NewDomainHandler())
	applicationApp.Handle(v1.NewBlackIPHandler())
	applicationApp.Handle(v1.NewIPHandler())
}
