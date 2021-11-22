package api

import (
	"bieshu-oa/app/model"
	"bieshu-oa/app/service"
	"bieshu-oa/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var UploadApi = uploadApi{}

type uploadApi struct{}

// File 文件上传
func (a *uploadApi) File(r *ghttp.Request) {

	var (
		req *model.UploadFileReq
	)

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	file := r.GetUploadFile("file")
	if file == nil {
		response.JsonExit(r, 1, g.I18n().Tf(r.Context(), `upload.file.not.exist`))
	}

	if data, err := service.UploadService.File(r.Context(), file, req.UploadFileInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", data)
	}
}
