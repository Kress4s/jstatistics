package tools

import (
	"fmt"
	"js_statistics/app/models"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/constant"
	"js_statistics/exception"

	"github.com/kataras/iris/v12"
)

// 异常返回
func ErrorResponse(ctx iris.Context, ex exception.Exception) {
	if ex == nil {
		ex = exception.New(response.ExceptionUnknown, "some error")
	}
	ctx.WriteString("occur error: " + ex.Error())
}

// 跳转管理全都关闭，默认空白页
func DefaultBlackCode(ctx iris.Context) {
	ctx.WriteString(fmt.Sprintf(constant.RedirectWindowsPage, constant.BlankCode))
}

// js规则之外的条件，设置的伪装内容
func BeyondRuleRedirect(ctx iris.Context, faker *vo.FakerResp, js *models.JsManage) {
	// var redirectInfo string
	if faker != nil {
		if !faker.Status {
			ScreenRedirect(ctx, constant.BlankCode)
		}
		var redirectInfo string
		switch faker.Type {
		//文本
		case 0:
			// text/html
			switch faker.ReqType {
			// text/html
			case 0:
				redirectInfo = fmt.Sprintf(constant.TextHtml, faker.Text)
			// text/plain
			case 1:
				redirectInfo = faker.Text
			// text/xml
			case 2:
				redirectInfo = fmt.Sprintf(constant.TextXml, faker.Text)
			// text/application
			case 3:
				redirectInfo = fmt.Sprintf(constant.ApplicationJson, faker.Text)
			}
		// 图片
		case 1:
			redirectInfo = GetMiniIoURL(faker.ObjID)
		// mp3
		case 2:
			redirectInfo = GetMiniIoMP3URL(faker.ObjID)
		// mp4
		case 3:
			redirectInfo = GetMiniIoURL(faker.ObjID)
		}
		// if redirectMode == 0 {
		// 	// ctx.WriteString(fmt.Sprintf(constant.RedirectWindowsPage, redirectInfo))
		// 	DirectWindowsRedirect(ctx, redirectInfo)
		// } else {
		// 	// ctx.WriteString(fmt.Sprintf(constant.RedirectTopPage, redirectInfo))
		// 	DirectTopRedirect(ctx, redirectInfo)
		// }
		switch js.RedirectMode {
		case constant.Direct:
			if js.RedirectCode == 0 {
				DirectWindowsRedirect(ctx, redirectInfo)
			} else {
				DirectTopRedirect(ctx, redirectInfo)
			}
		case constant.Nested:
			NestedRedirect(ctx, redirectInfo)
		case constant.Screen:
			ScreenRedirect(ctx, redirectInfo)
		default:
			// id 为动态参数
			// HrefRedirect(ctx, redirectInfo+"/"+strings.ReplaceAll(js.HrefID, ",", "/"))
			//TODO 不使用Href跳转
			ScreenRedirect(ctx, redirectInfo)
		}
	} else {
		// if redirectMode == 0 {
		// 未设置伪装内容 直接默认新标签页跳转
		// DirectWindowsRedirect(ctx, constant.BlankCode)
		ScreenRedirect(ctx, constant.BlankCode)
		// } else {
		// DirectTopRedirect(ctx, constant.BlankCode)
		// }
	}
}

func DirectWindowsRedirect(ctx iris.Context, redirect string) {
	ctx.WriteString(fmt.Sprintf(constant.RedirectWindowsPage, redirect))
}

func DirectTopRedirect(ctx iris.Context, redirect string) {
	ctx.WriteString(fmt.Sprintf(constant.RedirectTopPage, redirect))
}

func NestedRedirect(ctx iris.Context, redirect string) {
	ctx.WriteString(fmt.Sprintf(constant.NestingRedirect, redirect))
}

func ScreenRedirect(ctx iris.Context, redirect string) {
	ctx.WriteString(fmt.Sprintf(constant.ScreenRedirect, redirect))
}

func HrefRedirect(ctx iris.Context, redirect string) {
	ctx.WriteString(fmt.Sprintf(constant.HrefRedirect, redirect))
}
