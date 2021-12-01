package handlers

import (
	"js_statistics/app/vo"
	"js_statistics/constant"
	"js_statistics/exception"

	constants "js_statistics/constant"

	"js_statistics/app/response"
	"js_statistics/config"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type BaseHandler struct {
	UserID   int64
	UserName string
}

func (bh *BaseHandler) BeginRequest(ctx iris.Context) {
	token := ctx.Values().Get("jwt").(*jwt.Token)
	userInfo := token.Claims.(jwt.MapClaims)
	id := userInfo["user_id"].(float64)
	bh.UserID = int64(id)
	bh.UserName = userInfo["user_name"].(string)
}

func (bh *BaseHandler) EndRequest(ctx iris.Context) {}

func GetPageInfo(ctx iris.Context) (*vo.PageInfo, exception.Exception) {
	var page, pageSize int
	var err error
	maxPageSize := config.GetConfig().Server.MaxPageSize
	textSearch := ctx.URLParam(constants.TextSearch)
	switch {
	case ctx.URLParamExists(constants.Page) && ctx.URLParamExists(constants.PageSize):
		page, err = ctx.URLParamInt(constants.Page)
		if err != nil || page < 1 {
			page = 1
		}
		pageSize, err = ctx.URLParamInt(constants.PageSize)
		if err != nil || pageSize > maxPageSize {
			pageSize = maxPageSize
		}

	case !ctx.URLParamExists(constants.Page) && !ctx.URLParamExists(constants.PageSize):
		page = 1
		pageSize = maxPageSize

	default:
		return nil, exception.New(response.ExceptionMissingPageOrPageSize, "missing page or page_size")
	}

	return &vo.PageInfo{
		Page:     page,
		PageSize: pageSize,
		Keywords: textSearch,
	}, nil
}

func GetJSFilterParam(ctx iris.Context) (*vo.JSFilterParams, exception.Exception) {
	var pid, cid, jsID int64
	var err error
	if !ctx.URLParamExists(constant.PrimaryID) {
		return nil, exception.New(response.ExceptionInvalidRequestParameters, "pid not be null")
	}
	pid, err = ctx.URLParamInt64(constant.PrimaryID)
	if err != nil {
		pid = 0
	}
	if ctx.URLParamExists(constant.CategoryID) {
		cid, err = ctx.URLParamInt64(constant.CategoryID)
		if err != nil {
			cid = 0
		}
	}
	if ctx.URLParamExists(constant.JsID) {
		jsID, err = ctx.URLParamInt64(constant.JsID)
		if err != nil {
			jsID = 0
		}
	}
	return &vo.JSFilterParams{
		PrimaryID:  pid,
		CategoryID: cid,
		JsID:       jsID,
	}, nil
}

func GetTimeScopeParam(ctx iris.Context) (string, string, exception.Exception) {
	if !ctx.URLParamExists(constant.BeginAt) || !ctx.URLParamExists(constant.EndAt) {
		return "", "", exception.New(response.ExceptionMissingParameters, "begin_at or end_at cant be null")
	}
	return ctx.URLParam(constant.BeginAt), ctx.URLParam(constant.EndAt), nil
}
