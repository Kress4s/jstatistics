package v1

import (
	"js_statistics/app/handlers"
	"js_statistics/app/middlewares"
	"js_statistics/app/response"
	"js_statistics/app/service"
	"js_statistics/app/vo"
	"js_statistics/constant"
	"js_statistics/exception"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type FakerHandler struct {
	handlers.BaseHandler
	Svc service.FakerService
}

func NewFakerHandler() *FakerHandler {
	return &FakerHandler{
		Svc: service.GetFakerService(),
	}
}

// Create godoc
// @Summary 创建伪装内容(注意当 /application/faker/js/{js_id} 接口返回状态码404，才调用；否则Update 接口)
// @Description 创建伪装内容
// @Tags 应用管理 - 伪装内容
// @Param parameters body vo.FakerReq true "FakerReq"
// @Success 200  "创建伪装内容成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/faker [post]
func (fh *FakerHandler) Create(ctx iris.Context) mvc.Result {
	ip := &vo.FakerReq{}
	if err := ctx.ReadJSON(ip); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := fh.Svc.Create(fh.UserName, ip)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 查询伪装内容
// @Description 查询伪装内容信息
// @Tags 应用管理 - 伪装内容
// @Param id path string true "伪装内容id"
// @Success 200 {object} vo.FakerResp "查询伪装内容成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/faker/{id} [get]
func (fh *FakerHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := fh.Svc.Get(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// GetByJsID godoc
// @Summary 查询js下的伪装内容
// @Description 查询js下伪装内容信息
// @Tags 应用管理 - 伪装内容
// @Param js_id path string true "伪装内容的js ID"
// @Success 200 {object} vo.FakerResp "查询伪装内容成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/faker/js/{js_id} [get]
func (fh *FakerHandler) GetByJsID(ctx iris.Context) mvc.Result {
	jsID, err := ctx.Params().GetInt64(constant.JsID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := fh.Svc.GetByJsID(jsID)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 修改伪装内容(注意当 /application/faker/js/{js_id} 接口返回状态码200，才调用；否则Create 接口)
// @Description 修改伪装内容信息
// @Tags 应用管理 - 伪装内容
// @Param id path string true "伪装内容id"
// @Param parameters body vo.FakerUpdateReq true "FakerUpdateReq"
// @Success 200 "修改伪装内容成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/faker/{id} [put]
func (fh *FakerHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	req := &vo.FakerUpdateReq{}
	if err := ctx.ReadJSON(req); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := fh.Svc.Update(fh.UserName, id, req)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (fh *FakerHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/faker", "Create", middlewares.RecordSystemLog("Create", "", "创建伪装内容成功"))
	b.Handle(iris.MethodGet, "/faker/{id:string}", "Get")
	b.Handle(iris.MethodGet, "/faker/js/{js_id:string}", "GetByJsID")
	b.Handle(iris.MethodPut, "/faker/{id:string}", "Update", middlewares.RecordSystemLog("Update", "id", "更新伪装内容成功"))
}
