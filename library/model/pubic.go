package model

// PubicGetListPageReq 列表分页
type PubicGetListPageReq struct {
	Page  int `d:"1"  v:"min:0"`  // 分页号码
	Limit int `d:"10" v:"max:50"` // 分页数量，最大50
}

// PubicGetListPageInput 列表分页
type PubicGetListPageInput struct {
	Page  int // 分页号码
	Limit int // 分页数量，最大50
}
