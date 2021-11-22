package model

import (
	libraryModel "bieshu-oa/library/model"
)

type GroupIndexReq struct {
	GroupIndexInput
	libraryModel.PubicGetListPageReq
}

type GroupIndexInput struct {
	libraryModel.PubicGetListPageInput
}

type GroupIndexOutput struct {
	List []GroupList `json:"list"` // 列表
}

type GroupList struct {
	Id        string `json:"id"`         // ID
	Name      string `json:"name"`       // 组织架构名称
	Pid       int    `json:"pid"`        // 组织架构上级ID
	CreatedId string `json:"created_id"` // 添加人
	CreatedAt string `json:"created_at"` // 添加时间
	UpdatedId string `json:"updated_id"` // 最近更新人
	UpdatedAt string `json:"updated_at"` // 最近更新时间
}

type GroupStoreReq struct {
	GroupStoreInput
	Name string `v:"required"` // 组织架构名称
	Pid  int    `d:"0"`        // 组织架构上级ID
}

type GroupStoreInput struct {
	Name string // 组织架构名称
	Pid  int    // 组织架构上级ID
	Link string // 组织架构路径
}

type GroupUpdateReq struct {
	GroupUpdateInput
	Id   int    `v:"required|min:0|integer"` // 组织架构ID
	Name string `v:"required"`               // 组织架构名称
	Pid  int    `d:"0"`                      // 组织架构上级ID
}

type GroupUpdateInput struct {
	Id   int    // 组织架构ID
	Name string // 组织架构名称
	Pid  int    // 组织架构上级ID
	Link string // 组织架构路径
}

type GroupDeleteReq struct {
	Id int `v:"required|min:0|integer"` // 组织架构ID
}
