package model

type UploadFileOutput struct {
	Url string `json:"url"` // 文件地址
}

type UploadFileReq struct {
	UploadFileInput
	Type string `v:"required"` // 上传文件的类型
}

type UploadFileInput struct {
	Type string
}
