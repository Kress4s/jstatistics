package app

import (
	"fmt"
	"js_statistics/app/routes"
	"js_statistics/config"
	"log"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

// type Application struct {
// 	Server http.Server
// 	Routes *mux.Router
// }

func Run(port int) {
	if err := newApp().Run(iris.Addr(fmt.Sprintf("0.0.0.0:%d", port))); err != nil {
		log.Fatal("Web server run failed, err is ", err.Error())
	}
}

// newApp
func newApp() *iris.Application {
	cfg := config.GetConfig()
	app := iris.New()
	iris.WithOptimizations(app)
	// app.Use(middlewares.Recover())
	// if cfg.DebugMode && log.GetLevel() > log.InfoLevel {
	// 	app.Use(middlewares.IrisLog())
	// }
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
