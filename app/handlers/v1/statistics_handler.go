package v1

import (
	"js_statistics/app/handlers"
	"js_statistics/app/service"

	"github.com/kataras/iris/v12"
)

type StatisticHandler struct {
	handlers.BaseHandler
	WIPSvc service.WhiteIPService
	Svc    service.StcService
}

func NewStatisticHandler() *StatisticHandler {
	return &StatisticHandler{
		WIPSvc: service.GetWhiteIPService(),
		Svc:    service.GetStcService(),
	}
}

// Create godoc
// @Summary js链接(开发中, 暂不可调)
// @Description js链接
// @Tags JS链接
// @Param sign path string true "js唯一字符串"
// @Router / [get]
func (sh *StatisticHandler) FilterJS(ctx iris.Context) {
	sh.Svc.ProcessJsRequest(ctx)
}
