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

func RegisterJSRoutes(app *iris.Application) {
	// cfg := config.GetConfig()
	app.Get("/{sign:string}", v1.NewStatisticHandler().FilterJS)
}

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
	permissionApp.Handle(v1.NewSyslogHandler())

	// 应用管理
	applicationParty := party.Party("/application")
	applicationApp := mvc.New(applicationParty)
	applicationApp.Handle(v1.NewDomainHandler())
	applicationApp.Handle(v1.NewBlackIPHandler())
	applicationApp.Handle(v1.NewIPHandler())
	applicationApp.Handle(v1.NewCdnHandler())
	applicationApp.Handle(v1.NewJsPrimaryHandler())
	applicationApp.Handle(v1.NewJsCategoryHandler())
	applicationApp.Handle(v1.NewJsManageHandler())
	applicationApp.Handle(v1.NewRedirectManageHandler())
	applicationApp.Handle(v1.NewFakerHandler())

	// 主页
	homeParty := party.Party("/home")
	homeApp := mvc.New(homeParty)
	homeApp.Handle(v1.NewHomeHandler())

	//数据统计
	analysisParty := party.Party("/analysis")
	analysisApp := mvc.New(analysisParty)
	analysisApp.Handle(v1.NewDataAnalysisHandler())
	analysisApp.Handle(v1.NewFromAnalysisHandler())

	// 文件上传
	objectParty := party.Party("/faker")
	objectApp := mvc.New(objectParty)
	objectApp.Handle(v1.NewObjectHandler())

	// 文件访问 无权限
	// noAuthParty := app.Party("/object/v1")
	// app.Get("/{sign:string}", v1.NewStatisticHandler().FilterJS)
	app.Get("/object/{id:string}", v1.NewObjectHandler().Get)

}
