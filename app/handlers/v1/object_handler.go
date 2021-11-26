package v1

import (
	"js_statistics/app/handlers"
	"js_statistics/app/response"
	"js_statistics/app/service"
	"js_statistics/app/vo"
	"js_statistics/constant"
	"js_statistics/exception"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ObjectHandler struct {
	Service service.ObjectService
	handlers.BaseHandler
}

// NewObjectHandler ObjectHandler
func NewObjectHandler() *ObjectHandler {
	return &ObjectHandler{
		Service: service.GetObjectService(),
	}
}

// Upload godoc
// @Summary 存储对象
// @Description 存储对象并返回对象的id
// @Tags 应用管理 - 伪装内容 - 文件对象
// @Accept mpfd
// @Produce json
// @Param uploadfile formData file true "文件"
// @Success 201 {object} types.UUID "响应成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/faker/object [post]
func (oh *ObjectHandler) Upload(ctx iris.Context) mvc.Result {
	file, info, err := ctx.FormFile(constant.File)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}

	id, ex := oh.Service.UploadFromReader(oh.UserName, info.Filename, info.Size, file)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(vo.UUID{
		ID: id,
	})
}

// Get godoc
// @Summary 获取对象
// @Description 获取对象
// @Tags 应用管理 - 伪装内容 - 文件对象
// @Param id path string true "对象id"
// @Success 200 {string} byte "获取文件成功"
// @Failure 500 {string} byte "服务器内部错误"
// @Router /object/{id} [get]
func (oh *ObjectHandler) Get(ctx iris.Context) {
	id := ctx.Params().Get(constant.ID)

	obj, ex := oh.Service.Download(id)
	if ex != nil {
		ctx.ResponseWriter().Write([]byte(ex.Error()))
	}
	ctx.ResponseWriter().Write([]byte(obj.Content))
}

// BeforeActivation 初始化路由
func (oh *ObjectHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/object", "Upload")
}
