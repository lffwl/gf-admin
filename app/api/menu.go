package api

import (
	"bieshu-oa/app/model"
	"bieshu-oa/app/service"
	"bieshu-oa/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var MenuApi = menuApi{}

type menuApi struct{}

// Index 组织架构列表
func (a *menuApi) Index(r *ghttp.Request) {

	var (
		req *model.MenuIndexReq
	)

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if list, err := service.MenuService.Index(r.Context(), req.MenuIndexInput, true); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", list)
	}
}

// Store 新增
func (a *menuApi) Store(r *ghttp.Request) {

	var req *model.MenuStoreReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.MenuService.Store(r.Context(), req.MenuStoreInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// Update 更新
func (a *menuApi) Update(r *ghttp.Request) {

	var req *model.MenuUpdateReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.MenuService.Update(r.Context(), req.MenuUpdateInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// Delete 删除
func (a *menuApi) Delete(r *ghttp.Request) {

	var req *model.MenuDeleteReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.MenuService.Delete(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}
