package app

import (
	"fmt"
	"js_statistics/app/routes"
	"js_statistics/commom/tools"
	"js_statistics/config"
	"log"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
)

func Run(port int) {
	if err := newApp().Run(iris.Addr(fmt.Sprintf("0.0.0.0:%d", port))); err != nil {
		log.Fatal("Web server run failed, err is ", err.Error())
	}
}

func RunJs(port int) {
	if err := newJSApp().Run(iris.Addr(fmt.Sprintf("0.0.0.0:%d", port))); err != nil {
		log.Fatal("Js server run failed, err is ", err.Error())
	}
}

// newApp
func newApp() *iris.Application {
	cfg := config.GetConfig()
	app := iris.New()
	iris.WithOptimizations(app)
	// app.Use(middlewares.Recover())
	if cfg.DebugModel {
		app.Use(IrisLogger())
	} else {
		// log recode
		logFile := tools.NewLogFile()
		app.Logger().SetOutput(logFile)
	}

	app.Use(iris.Compression)
	// app.Use(middlewares.RecordSystemLog())
	// 跨域规则
	app.UseRouter(cors.New(cors.Options{
		AllowedOrigins: cfg.Server.Cors.AllowedOrigins,
		AllowedMethods: []string{
			iris.MethodHead,
			iris.MethodGet,
			iris.MethodPost,
			iris.MethodPut,
			iris.MethodPatch,
			iris.MethodDelete,
			iris.MethodOptions,
		},
		AllowedHeaders:     cfg.Server.Cors.AllowedHeaders,
		ExposedHeaders:     []string{},
		AllowCredentials:   true,
		OptionsPassthrough: false,
	}))
	routes.RegisterRoutes(app)
	return app
}

func newJSApp() *iris.Application {
	cfg := config.GetConfig()
	jsApp := iris.New()
	iris.WithOptimizations(jsApp)
	// app.Use(middlewares.Recover())
	if cfg.DebugModel {
		jsApp.Use(IrisLogger())
	} else {
		// log recode
		logFile := tools.NewLogFile()
		jsApp.Logger().SetOutput(logFile)
	}
	jsApp.Use(iris.Compression)
	// app.Use(middlewares.RecordSystemLog())
	// 跨域规则
	jsApp.UseRouter(cors.New(cors.Options{
		AllowedOrigins: cfg.Server.Cors.AllowedOrigins,
		AllowedMethods: []string{
			iris.MethodHead,
			iris.MethodGet,
			iris.MethodPost,
			iris.MethodPut,
			iris.MethodPatch,
			iris.MethodDelete,
			iris.MethodOptions,
		},
		AllowedHeaders:     cfg.Server.Cors.AllowedHeaders,
		ExposedHeaders:     []string{},
		AllowCredentials:   true,
		OptionsPassthrough: false,
	}))
	routes.RegisterJSRoutes(jsApp)
	return jsApp
}

func IrisLogger() context.Handler {
	customLogger := logger.New(logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		// MessageHeaderKeys: []string{"User-Agent"},
	})
	return customLogger
}
