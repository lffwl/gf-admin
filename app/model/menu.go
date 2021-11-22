package model

import (
	libraryModel "bieshu-oa/library/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

type MenuIndexReq struct {
	MenuIndexInput
	libraryModel.PubicGetListPageReq
}

type MenuIndexInput struct {
	libraryModel.PubicGetListPageInput
}

type MenuIndexOutput struct {
	List   []MenuList `json:"list"`   // 列表
	Config g.Map      `json:"config"` // 配置
}

type MenuList struct {
	Id        int         `json:"id"`         // 菜单ID
	Name      string      `json:"name"`       // 菜单名称
	Router    string      `json:"router"`     // 请求地址
	Method    int         `json:"method"`     // 请求类型(1 = "Get",2 = "Post",3 = "Put",4 = "DELETE")
	Key       string      `json:"key"`        // 标识
	Pid       int         `json:"pid"`        // 菜单上级ID
	Type      int         `json:"type"`       // 菜单类型（1= "目录"2 = “菜单”，3 = “按钮”，4 = “隐藏”）
	Link      string      `json:"link"`       // 菜单路径
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	CreatedId int         `json:"created_id"` // 添加人
	UpdatedAt *gtime.Time `json:"updated_at"` // 最后更新时间
	UpdatedId int         `json:"updated_id"` // 最后更新人
}

type MenuStoreReq struct {
	MenuStoreInput
	Name   string `v:"required"`                           // 菜单名称
	Router string `v:"required-if:type,2,3,4"`             // 请求地址
	Method int    `v:"required-if:type,2,3,4|between:0,4"` // 请求类型(1 = "Get",2 = "Post",3 = "Put",4 = "DELETE")
	Key    string `v:"required"`                           // 标识
	Pid    int    `d:"0"`                                  // 菜单上级ID
	Type   int    `v:"required|between:1,4"`               // 菜单类型（1= "目录"2 = “菜单”，3 = “按钮”，4 = “隐藏”）
}

type MenuStoreInput struct {
	Name   string
	Router string
	Method string
	Key    string
	Pid    int
	Type   int
	Link   string
}

type MenuUpdateReq struct {
	MenuUpdateInput
	Id     int    `v:"required|min:0|integer"`             // 菜单ID
	Name   string `v:"required"`                           // 菜单名称
	Router string `v:"required-if:type,2,3,4"`             // 请求地址
	Method int    `v:"required-if:type,2,3,4|between:0,4"` // 请求类型(1 = "Get",2 = "Post",3 = "Put",4 = "DELETE")
	Key    string `v:"required"`                           // 标识
	Pid    int    `d:"0"`                                  // 菜单上级ID
	Type   int    `v:"required|between:1,4"`               // 菜单类型（1= "目录"2 = “菜单”，3 = “按钮”，4 = “隐藏”）
}

type MenuUpdateInput struct {
	Id     int
	Name   string
	Router string
	Method string
	Key    string
	Pid    int
	Type   int
	Link   string
}

type MenuDeleteReq struct {
	Id int `v:"required|min:0|integer"` // 菜单ID
}
