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

type JspHandler struct {
	handlers.BaseHandler
	Svc service.JspService
}

func NewJsPrimaryHandler() *JspHandler {
	return &JspHandler{
		Svc: service.GetJspService(),
	}
}

// Create godoc
// @Summary 创建js主分类
// @Description 创建js主分类管理
// @Tags 应用管理 - js主分类
// @Param parameters body vo.JsPrimaryReq true "JsPrimaryReq"
// @Success 200  "创建js主分类成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_primary [post]
func (ch *JspHandler) Create(ctx iris.Context) mvc.Result {
	cdn := &vo.JsPrimaryReq{}
	if err := ctx.ReadJSON(cdn); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := ch.Svc.Create(ch.UserName, cdn)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 查询js主分类列表
// @Description 查询js主分类列表
// @Tags 应用管理 - js主分类
// @Success 200 {array}  vo.JsPrimaryResp "查询js主分类列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_primaries [get]
func (ch *JspHandler) List(ctx iris.Context) mvc.Result {
	resp, ex := ch.Svc.List()
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 查询js主分类
// @Description 查询js主分类信息
// @Tags 应用管理 - js主分类
// @Param id path string true "js主分类id"
// @Success 200 {object} vo.JsPrimaryResp "查询js主分类信息成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_primary/{id} [get]
func (ch *JspHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := ch.Svc.Get(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 修改js主分类
// @Description 修改js主分类信息
// @Tags 应用管理 - js主分类
// @Param id path string true "js主分类id"
// @Param parameters body vo.JsPrimaryUpdateReq true "JsPrimaryUpdateReq"
// @Success 200 "修改js主分类成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_primary/{id} [put]
func (ch *JspHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	param := &vo.JsPrimaryUpdateReq{}
	if err := ctx.ReadJSON(param); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := ch.Svc.Update(ch.UserName, id, param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 删除js主分类
// @Description 删除js主分类信息
// @Tags 应用管理 - js主分类
// @Param id path string true "js主分类id"
// @Success 200 "删除js主分类成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_primary/{id} [delete]
func (ch *JspHandler) Delete(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ex := ch.Svc.Delete(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// GetAllsCategoryTree godoc
// @Summary 获取js主分类下的分类树(只到分类)
// @Description 获取js主分类下的分类树
// @Tags 应用管理 - js主分类
// @Success 200 {array} vo.Primaries "获取js主分类下的分类树成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_primary/category/tree [get]
func (ch *JspHandler) GetAllsCategoryTree(ctx iris.Context) mvc.Result {
	resp, ex := ch.Svc.GetAllsCategoryTree()
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// BeforeActivation 初始化路由
func (ch *JspHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/js_primary", "Create", middlewares.RecordSystemLog("Create", "", "创建JS分类成功"))
	b.Handle(iris.MethodGet, "/js_primaries", "List")
	b.Handle(iris.MethodGet, "/js_primary/{id:string}", "Get")
	b.Handle(iris.MethodGet, "/js_primary/category/tree", "GetAllsCategoryTree")
	b.Handle(iris.MethodPut, "/js_primary/{id:string}", "Update", middlewares.RecordSystemLog("Update", "id", "更新JS分类成功"))
	b.Handle(iris.MethodDelete, "/js_primary/{id:string}", "Delete", middlewares.RecordSystemLog("Delete", "id", "删除JS分类成功"))
}
