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

type JscHandler struct {
	handlers.BaseHandler
	Svc service.JscService
}

func NewJsCategoryHandler() *JscHandler {
	return &JscHandler{
		Svc: service.GetJscService(),
	}
}

// Create godoc
// @Summary 创建js分类
// @Description 创建js分类管理
// @Tags 应用管理 - js分类
// @Param parameters body vo.JsCategoryReq true "JsCategoryReq"
// @Success 200  "创建js分类成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_category [post]
func (jh *JscHandler) Create(ctx iris.Context) mvc.Result {
	cdn := &vo.JsCategoryReq{}
	if err := ctx.ReadJSON(cdn); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := jh.Svc.Create(jh.UserName, cdn)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// ListByPrimaryID godoc
// @Summary 查询js分类列表通过js主分类ID
// @Description 查询js分类列表通过js主分类ID
// @Tags 应用管理 - js分类
// @Param pid path string true "js主分类id"
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Success 200 {object} vo.DataPagination{data=[]vo.JsCategoryResp} "查询js分类列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_categories/primary/{pid} [get]
func (jh *JscHandler) ListByPrimaryID(ctx iris.Context) mvc.Result {
	pid, err := ctx.Params().GetInt64(constant.PrimaryID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	params, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := jh.Svc.ListByPrimaryID(params, pid, jh.UserID)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Get godoc
// @Summary 查询js分类
// @Description 查询js分类信息
// @Tags 应用管理 - js分类
// @Param id path string true "js分类id"
// @Success 200 {object} vo.JsCategoryResp "查询域名成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_category/{id} [get]
func (jh *JscHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := jh.Svc.Get(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Update godoc
// @Summary 修改js分类
// @Description 修改js分类信息
// @Tags 应用管理 - js分类
// @Param id path string true "js分类id"
// @Param parameters body vo.JsCategoryUpdateReq true "JsCategoryUpdateReq"
// @Success 200 "修改js分类成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_category/{id} [put]
func (jh *JscHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	param := &vo.JsCategoryUpdateReq{}
	if err := ctx.ReadJSON(param); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := jh.Svc.Update(jh.UserName, id, param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Delete godoc
// @Summary 删除js分类
// @Description 删除js分类信息
// @Tags 应用管理 - js分类
// @Param id path string true "js分类id"
// @Success 200 "删除js分类成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_category/{id} [delete]
func (jh *JscHandler) Delete(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ex := jh.Svc.Delete(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// MultiDelete godoc
// @Summary 批量删除js分类
// @Description 批量删除js分类信息
// @Tags 应用管理 - js分类
// @Param ids query string true "js分类ids, `,` 连接"
// @Success 200 "批量删除js分类管理成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_category/multi [delete]
func (jh *JscHandler) MultiDelete(ctx iris.Context) mvc.Result {
	ids := ctx.URLParam(constant.IDS)
	ex := jh.Svc.MultiDelete(ids)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// ListAllByvPrimaryID godoc
// @Summary 查询js所有无分页分类通过js主分类ID(js管理table表头使用)
// @Description 查询js所有无分页分类通过js主分类ID
// @Tags 应用管理 - js分类
// @Param pid path string true "js主分类id"
// @Success 200 {array} vo.JsCategoryResp "查询js分类列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/all/js_categories/primary/{pid} [get]
func (jh *JscHandler) ListAllByPrimaryID(ctx iris.Context) mvc.Result {
	pid, err := ctx.Params().GetInt64(constant.PrimaryID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	res, ex := jh.Svc.ListAllByPrimaryID(pid, jh.UserID)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(res)
}

// BeforeActivation 初始化路由
func (jh *JscHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/js_category", "Create", middlewares.RecordSystemLog("Create", "", "创建js分类成功"))
	b.Handle(iris.MethodGet, "/js_categories/primary/{pid:string}", "ListByPrimaryID")
	b.Handle(iris.MethodGet, "/js_category/{id:string}", "Get")
	b.Handle(iris.MethodPut, "/js_category/{id:string}", "Update", middlewares.RecordSystemLog("Update", "id", "更新js分类成功"))
	b.Handle(iris.MethodDelete, "/js_category/{id:string}", "Delete", middlewares.RecordSystemLog("Delete", "id", "删除js分类成功"))
	b.Handle(iris.MethodDelete, "/js_category/multi", "MultiDelete", middlewares.RecordSystemLog("MultiDelete", "ids", "批量删除js分类成功"))
	b.Handle(iris.MethodGet, "/all/js_categories/primary/{pid:string}", "ListAllByPrimaryID")
}
