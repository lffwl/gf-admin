package api

import (
	"bieshu-oa/app/model"
	"bieshu-oa/app/service"
	"bieshu-oa/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var RoleApi = roleApi{}

type roleApi struct{}

// Index 组织架构列表
func (a *roleApi) Index(r *ghttp.Request) {

	var (
		req *model.RoleIndexReq
	)

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if list, err := service.RoleService.Index(r.Context(), req.RoleIndexInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", list)
	}
}

// Store 新增
func (a *roleApi) Store(r *ghttp.Request) {

	var req *model.RoleStoreReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.RoleService.Store(r.Context(), req.RoleStoreInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// Update 更新
func (a *roleApi) Update(r *ghttp.Request) {

	var req *model.RoleUpdateReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.RoleService.Update(r.Context(), req.RoleUpdateInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// Delete 删除
func (a *roleApi) Delete(r *ghttp.Request) {

	var req *model.RoleDeleteReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.RoleService.Delete(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// Show 详情
func (a *roleApi) Show(r *ghttp.Request) {

	var req *model.RoleDeleteReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if info, err := service.RoleService.Show(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", info)
	}
}
