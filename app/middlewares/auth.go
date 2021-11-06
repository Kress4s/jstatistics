package middlewares

import (
	"js_statistics/constant"

	"github.com/iris-contrib/middleware/jwt"
)

func Auth() *jwt.Middleware {
	return jwt.New(jwt.Config{
		// Extractor属性可以选择从什么地方获取jwt进行验证，默认从http请求的header中的Authorization字段提取，也可指定为请求参数中的某个字段

		// 从请求参数token中提取
		// Extractor: jwt.FromParameter("token"),

		// 从请求头的Authorization字段中提取，这个是默认值
		Extractor: jwt.FromAuthHeader,

		// 设置一个函数返回秘钥，关键在于return []byte("这里设置秘钥")
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(constant.Salt), nil
		},

		// 设置一个加密方法
		SigningMethod: jwt.SigningMethodHS256,
	})
}
