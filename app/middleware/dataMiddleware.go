package middleware

import (
	"bieshu-oa/app/define"
	"bieshu-oa/app/service"
	"bieshu-oa/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var DataMiddleware = dataMiddleware{}

type dataMiddleware struct{}

// GetDataAdminId 获取可以操作的数据的管理员ID
func (s *dataMiddleware) GetDataAdminId(r *ghttp.Request) {
	// 当前的adminId
	currentAdminId := gconv.Int(r.Context().Value(define.CurrentAdminId))
	// 获取可以管理的AdminId
	if admins, err := service.AdminService.GetAdminDataAdminId(r.Context(), currentAdminId); err != nil {
		response.JsonExit(r, 403, g.I18n().Tf(r.Context(), `auth.get.data.error`))
	} else {
		r.SetCtxVar(define.CurrentAdminData, admins)
	}
	r.Middleware.Next()
}
