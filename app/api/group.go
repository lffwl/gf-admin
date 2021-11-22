package api

import (
	"bieshu-oa/app/model"
	"bieshu-oa/app/service"
	"bieshu-oa/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var GroupApi = groupApi{}

type groupApi struct{}

// Index 组织架构列表
func (a *groupApi) Index(r *ghttp.Request) {

	var (
		req *model.GroupIndexReq
	)

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if list, err := service.GroupService.Index(r.Context(), req.GroupIndexInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", list)
	}
}

// Store 新增
func (a *groupApi) Store(r *ghttp.Request) {

	var req *model.GroupStoreReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.GroupService.Store(r.Context(), req.GroupStoreInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// Update 更新
func (a *groupApi) Update(r *ghttp.Request) {

	var req *model.GroupUpdateReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.GroupService.Update(r.Context(), req.GroupUpdateInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// Delete 删除
func (a *groupApi) Delete(r *ghttp.Request) {

	var req *model.GroupDeleteReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.GroupService.Delete(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}
