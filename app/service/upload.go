package service

import (
	"bieshu-oa/app/model"
	"bieshu-oa/library/helper"
	"context"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

// PublicPath 隐私文件前缀
const PublicPath = "/public"

// PrivatePath 非隐私文件前缀
const PrivatePath = "/file"

var UploadService = uploadService{}

type uploadService struct{}

// File 文件上传
func (s *uploadService) File(
	ctx context.Context,
	file *ghttp.UploadFile,
	input model.UploadFileInput,
) (output *model.UploadFileOutput, err error) {

	output = &model.UploadFileOutput{}

	// 获取配置
	config := g.Cfg("upload").GetMap(input.Type)
	if config == nil {
		return nil, errors.New(g.I18n().Tf(ctx, `upload.file.type.error`))
	}

	var (
		path string
		//是否隐私文件
		private = gconv.Bool(config["private"])
		// 文件夹
		prefix = gconv.String(config["path"])
		// 文件大小
		size = gconv.Int64(config["size"])
		// 文件类型
		fileType = gconv.Strings(config["type"])
	)

	// 验证大小是否超过
	if file.Size > size {
		return nil, errors.New(g.I18n().Tf(ctx, `upload.file.size.error`, helper.File.FormatFileSize(size)))
	}

	// 验证文件类型
	if helper.Array.In(file.Header.Get("Content-Type"), fileType) == false {
		return nil, errors.New(g.I18n().Tf(ctx, `upload.file.content.type.error`, strings.Join(fileType, ",")))
	}

	// 文件路径，是否隐私文件
	if private {
		path = g.Cfg().GetString("upload.path") + PrivatePath + prefix
	} else {
		path = g.Cfg().GetString("upload.path") + PublicPath + prefix
	}

	//文件上传
	if output.Url, err = file.Save(path, true); err != nil {
		return nil, err
	}

	output.Url = prefix + output.Url

	return
}
