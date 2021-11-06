package response

import (
	"fmt"
	"js_statistics/exception"
	"log"
	"mime"
	"path"
	"strings"

	"js_statistics/app/vo"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/v12/mvc"
)

func Error(ex exception.Exception, args ...string) mvc.Result {
	return mvc.Response{
		Code:        GetStatusCode(ex),
		ContentType: context.ContentJSONHeaderValue,
		Object: vo.Error{
			Code: ex.Type().Code(),
			Msg:  ex.Error(),
			Args: args,
		},
	}
}

// OK response ok to front
func OK() mvc.Result {
	return mvc.Response{
		Code: iris.StatusOK,
	}
}

// ID response id to front
func ID(id int64) mvc.Result {
	return mvc.Response{
		Code:        iris.StatusCreated,
		ContentType: context.ContentJSONHeaderValue,
		Object: vo.ID{
			ID: id,
		},
	}
}

// Object response object to front
func Object(filename string, content []byte) mvc.Result {
	return mvc.Response{
		Code: iris.StatusOK,
		Object: &object{
			Filename: filename,
			Content:  content,
		},
	}
}

// JSON response json to front
func JSON(v interface{}) mvc.Result {
	return mvc.Response{
		Code:        iris.StatusOK,
		ContentType: context.ContentJSONHeaderValue,
		Object:      v,
	}
}

// Text response text to front
func Text(v string) mvc.Result {
	return mvc.Response{
		Code: iris.StatusOK,
		Text: v,
	}
}

// HTML response html to front
func HTML(format string, args ...interface{}) mvc.Result {
	return mvc.Response{
		Code:        iris.StatusOK,
		ContentType: context.ContentHTMLHeaderValue,
		Content:     []byte(fmt.Sprintf(format, args...)),
	}
}

// Redirect send redirect url to front
func Redirect(path string) mvc.Result {
	return mvc.Response{
		Code: iris.StatusFound,
		Path: path,
	}
}

type object struct {
	Filename string
	Content  []byte
}

func (o *object) Dispatch(ctx *context.Context) {
	ctx.ResponseWriter().Header().
		Set(context.ContentTypeHeaderKey, mime.TypeByExtension(strings.ToLower(path.Ext(o.Filename))))
	ctx.ResponseWriter().Header().
		Set(context.ContentDispositionHeaderKey, fmt.Sprintf("attachment;filename=%s", o.Filename))
	_, err := ctx.Write(o.Content)
	if err != nil {
		log.Fatal("write content error:", err)
	}
}
