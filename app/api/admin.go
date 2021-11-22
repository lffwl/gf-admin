package api

import (
	"bieshu-oa/app/model"
	"bieshu-oa/app/service"
	"bieshu-oa/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var AdminApi = adminApi{}

type adminApi struct{}

// Index 列表
func (a *adminApi) Index(r *ghttp.Request) {

	var (
		req *model.AdminIndexReq
	)

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if list, err := service.AdminService.Index(r.Context(), req.AdminIndexInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", list)
	}
}

// Store 新增
func (a *adminApi) Store(r *ghttp.Request) {

	var req *model.AdminStoreReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.AdminService.Store(r.Context(), req.AdminStoreInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// Update 更新
func (a *adminApi) Update(r *ghttp.Request) {

	var req *model.AdminUpdateReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.AdminService.Update(r.Context(), req.AdminUpdateInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// Delete 删除
func (a *adminApi) Delete(r *ghttp.Request) {

	var req *model.AdminDeleteReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.AdminService.Delete(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// Show 详情
func (a *adminApi) Show(r *ghttp.Request) {

	var req *model.AdminDeleteReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if info, err := service.AdminService.Show(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", info)
	}
}
