package response

import (
	"js_statistics/exception"

	"github.com/kataras/iris/v12"
)

type Exception struct {
	code       int
	statusCode int
}

func (e Exception) Code() int {
	return e.code
}

func GetStatusCode(e exception.Exception) int {
	if e == nil {
		return iris.StatusOK
	}
	et, ok := e.Type().(*Exception)
	if !ok {
		return iris.StatusInternalServerError
	}
	return et.statusCode
}

var (
	ExceptionInvalidRequestBody       exception.Type = &Exception{code: 400001, statusCode: iris.StatusBadRequest}
	ExceptionMissingParameters        exception.Type = &Exception{code: 400002, statusCode: iris.StatusBadRequest}
	ExceptionInvalidRequestParameters exception.Type = &Exception{code: 400003, statusCode: iris.StatusBadRequest}
	ExceptionMissingPageOrPageSize    exception.Type = &Exception{code: 400005, statusCode: iris.StatusBadRequest}
	ExceptionInvalidUserPassword      exception.Type = &Exception{code: 400006, statusCode: iris.StatusBadRequest}
	ExceptionInvalidFile              exception.Type = &Exception{code: 400009, statusCode: iris.StatusBadRequest}
	ExceptionNotConfigure             exception.Type = &Exception{code: 400011, statusCode: iris.StatusBadRequest}
	ExceptionNameDuplicate            exception.Type = &Exception{code: 400012, statusCode: iris.StatusBadRequest}
	ExceptionInvalidAccessToken       exception.Type = &Exception{code: 401001, statusCode: iris.StatusUnauthorized}
	ExceptionForbidden                exception.Type = &Exception{code: 403001, statusCode: iris.StatusForbidden}
	ExceptionRecordNotFound           exception.Type = &Exception{code: 404001, statusCode: iris.StatusNotFound}
	ExceptionUserClose                exception.Type = &Exception{code: 405001, statusCode: iris.StatusNotFound}
	ExceptionUnknown                  exception.Type = &Exception{code: 500000, statusCode: iris.StatusInternalServerError}
	ExceptionMarshalJSON              exception.Type = &Exception{code: 500001, statusCode: iris.StatusInternalServerError}
	ExceptionUnmarshalJSON            exception.Type = &Exception{code: 500002, statusCode: iris.StatusInternalServerError}
	ExceptionDatabase                 exception.Type = &Exception{code: 500003, statusCode: iris.StatusInternalServerError}
	ExceptionVo2Model                 exception.Type = &Exception{code: 500004, statusCode: iris.StatusInternalServerError}
	ExceptionModel2Vo                 exception.Type = &Exception{code: 500005, statusCode: iris.StatusInternalServerError}
	ExceptionGetUserInfo              exception.Type = &Exception{code: 500006, statusCode: iris.StatusInternalServerError}
	ExceptionDeleteObject             exception.Type = &Exception{code: 500007, statusCode: iris.StatusInternalServerError}
	ExceptionDownloadObject           exception.Type = &Exception{code: 500008, statusCode: iris.StatusInternalServerError}
	ExceptionUploadObject             exception.Type = &Exception{code: 500009, statusCode: iris.StatusInternalServerError}
	ExceptionGenerateID               exception.Type = &Exception{code: 500010, statusCode: iris.StatusInternalServerError}
	ExceptionVo2Map                   exception.Type = &Exception{code: 500016, statusCode: iris.StatusInternalServerError}
	ExceptionParseFloatError          exception.Type = &Exception{code: 500022, statusCode: iris.StatusInternalServerError}
	ExceptionUnzipBundleError         exception.Type = &Exception{code: 500023, statusCode: iris.StatusInternalServerError}
	ExceptionParseDate                exception.Type = &Exception{code: 500024, statusCode: iris.StatusInternalServerError}
	ExceptionParseStringToInt64Error  exception.Type = &Exception{code: 500025, statusCode: iris.StatusInternalServerError}
	ExceptionHttpRequestError         exception.Type = &Exception{code: 500026, statusCode: iris.StatusInternalServerError}
	ExceptionPraseIPLocationError     exception.Type = &Exception{code: 500027, statusCode: iris.StatusInternalServerError}
)
