package middleware

import (
	"bieshu-oa/app/define"
	"bieshu-oa/app/service"
	"bieshu-oa/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

var BaseMiddleware = baseMiddleware{}

type baseMiddleware struct{}

// SetLanguage 设置国际化语言支持
func (s *baseMiddleware) SetLanguage(r *ghttp.Request) {
	// language 参数设置返回语言
	language := r.GetHeader(define.LANGUAGE)
	// 默认设置 zh-CN
	if language == "" {
		language = "zh-CN"
	}
	g.I18n().SetLanguage(language)
	r.Middleware.Next()
}

// CORS 允许接口跨域请求
func (s *baseMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// CheckAdminRouter 检查是否有权限访问当前地址
func (s *baseMiddleware) CheckAdminRouter(r *ghttp.Request) {
	if service.AuthService.CheckAdminRouter(r.Context(), r.Router.Uri, r.Router.Method) == false {
		response.JsonExit(r, 403, g.I18n().Tf(r.Context(), `auth.check.router.error`))
	}
	r.Middleware.Next()
}

// Dome 演示站点
func (s *baseMiddleware) Dome(r *ghttp.Request) {

	// 获取配置是否演示站
	if g.Cfg().GetBool("server.dome") {
		if strings.ToUpper(r.Router.Method) != "GET" {
			response.JsonExit(r, 1, g.I18n().Tf(r.Context(), `dome.error`))
		}
	}

	r.Middleware.Next()
}
