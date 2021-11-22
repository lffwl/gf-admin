package model

import (
	libraryModel "bieshu-oa/library/model"
	"github.com/gogf/gf/frame/g"
)

type RoleIndexReq struct {
	RoleIndexInput
	libraryModel.PubicGetListPageReq
	Name string `json:"name"` // 角色名称
}

type RoleIndexInput struct {
	libraryModel.PubicGetListPageInput
	Name string
}

type RoleIndexOutput struct {
	List   []RoleList `json:"list"`   // 列表
	Page   int        `json:"page"`   // 分页码
	Limit  int        `json:"limit"`  // 分页数量
	Total  int        `json:"total"`  // 数据总数
	Config g.Map      `json:"config"` // 配置
}

type RoleList struct {
	Id        string `json:"id"`         // ID
	Name      string `json:"name"`       // 角色名称
	Dp        int    `json:"dp"`         // 数据权限（0-仅自己，1-所在组织架构和下级组织架构，2-所有人）
	CreatedId string `json:"created_id"` // 添加人
	CreatedAt string `json:"created_at"` // 添加时间
	UpdatedId string `json:"updated_id"` // 最近更新人
	UpdatedAt string `json:"updated_at"` // 最近更新时间
}

type RoleStoreReq struct {
	RoleStoreInput
	Name  string `v:"required"`          // 角色名称
	Menus string `v:"required"`          // 菜单集合
	Dp    int    `d:"1" v:"between:0,2"` // 数据权限（0-仅自己，1-所在组织架构和下级组织架构，2-所有人）
}

type RoleStoreInput struct {
	Name  string
	Menus string
	Dp    int
}

type RoleUpdateReq struct {
	RoleUpdateInput
	Id    int    `v:"required|min:0|integer"` // 角色ID
	Name  string `v:"required"`               // 角色名称
	Menus string `v:"required"`               // 菜单集合
	Dp    int    `d:"1" v:"between:0,2"`      // 数据权限（0-仅自己，1-所在组织架构和下级组织架构，2-所有人）
}

type RoleUpdateInput struct {
	Id    int
	Name  string
	Menus string
	Dp    int
}

type RoleDeleteReq struct {
	Id int `v:"required|min:0|integer"` // 角色ID
}

type RoleShowOutput struct {
	Id    string   `json:"id"`    // ID
	Name  string   `json:"name"`  // 角色名称
	Dp    int      `json:"dp"`    // 数据权限（0-仅自己，1-所在组织架构和下级组织架构，2-所有人）
	Menus []string `json:"menus"` // 角色权限
}
