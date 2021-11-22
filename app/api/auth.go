package api

import (
	"bieshu-oa/app/define"
	"bieshu-oa/app/model"
	"bieshu-oa/app/service"
	"bieshu-oa/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var AuthApi = authApi{}

type authApi struct{}

// Login 登录
func (a *authApi) Login(r *ghttp.Request) (string, interface{}) {

	var (
		req *model.AuthLoginReq
	)

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	data, err := service.AuthService.Login(r.Context(), req.AuthLoginInput)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	return gconv.String(data["username"]), data["id"]
}

// Info 详情
func (a *authApi) Info(r *ghttp.Request) {

	if data, err := service.AuthService.Info(r.Context(), gconv.Int(r.GetCtxVar(define.CurrentAdminId))); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", data)
	}

}
