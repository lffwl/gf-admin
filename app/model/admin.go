package model

import (
	libraryModel "bieshu-oa/library/model"
	"github.com/gogf/gf/frame/g"
)

type AdminIndexReq struct {
	AdminIndexInput
	libraryModel.PubicGetListPageReq
	UserName string ``                // 用户名
	RealName string ``                // 真实姓名
	Mobile   string ``                // 手机号码
	Email    string ``                // 邮箱
	Status   string `v:"between:0,1"` // 状态（1-开启，0-关闭）
}

type AdminIndexInput struct {
	libraryModel.PubicGetListPageInput
	UserName string
	RealName string
	Mobile   string
	Email    string
	Status   string
}

type AdminIndexOutput struct {
	List   []AdminList      `json:"list"`   // 列表
	Page   int              `json:"page"`   // 分页码
	Limit  int              `json:"limit"`  // 分页数量
	Total  int              `json:"total"`  // 数据总数
	Config g.Map            `json:"config"` // 配置
	Roles  []AdminRoleList  `json:"roles"`  // 角色列表
	Groups []AdminGroupList `json:"groups"` // 组织架构列表
}

type AdminRoleList struct {
	Id   string `json:"id"`   // ID
	Name string `json:"name"` // 角色名称
}

type AdminGroupList struct {
	Id   string `json:"id"`   // ID
	Name string `json:"name"` // 组织架构名称
	Pid  int    `json:"pid"`  // 组织架构名称
}

type AdminList struct {
	Id        int    `json:"id"`         // 管理员ID
	UserName  string `json:"user_name"`  // 用户名
	RealName  string `json:"real_name"`  // 真实姓名
	Mobile    string `json:"mobile"`     // 手机号码
	Email     string `json:"email"`      // 邮箱
	Status    int    `json:"status"`     // 状态（1-开启，0-关闭）
	Avatar    string `json:"avatar"`     // 头像
	GroupId   int    `json:"group_id"`   // 组织架构
	CreatedId string `json:"created_id"` // 添加人
	CreatedAt string `json:"created_at"` // 添加时间
	UpdatedId string `json:"updated_id"` // 最近更新人
	UpdatedAt string `json:"updated_at"` // 最近更新时间
}

type AdminShowOutput struct {
	Id       int      `json:"id"`        // 管理员ID
	UserName string   `json:"user_name"` // 用户名
	RealName string   `json:"real_name"` // 真实姓名
	Mobile   string   `json:"mobile"`    // 手机号码
	Email    string   `json:"email"`     // 邮箱
	Status   int      `json:"status"`    // 状态（1-开启，0-关闭）
	Avatar   string   `json:"avatar"`    // 头像
	GroupId  int      `json:"group_id"`  // 组织架构
	Roles    []string `json:"roles"`     // 角色
}

type AdminStoreReq struct {
	AdminStoreInput
	UserName string `v:"required"`             // 用户名
	RealName string `v:"required"`             // 真实姓名
	Password string `v:"password2"`            // 密码
	Mobile   string ``                         // 手机号码
	Email    string ``                         // 邮箱
	Status   int    `v:"required|between:0,1"` // 状态（1-开启，0-关闭）
	Avatar   string ``                         // 头像
	Roles    string `v:"required"`             // 角色集合
	GroupId  int    `v:"required"`             // 组织架构
}

type AdminStoreInput struct {
	UserName string
	RealName string
	Password string
	Mobile   string
	Email    string
	Status   int
	Avatar   string
	Roles    string
	GroupId  int
}

type AdminUpdateReq struct {
	AdminUpdateInput
	Id       int    `v:"required|min:0|integer"` // 管理员ID
	UserName string `v:"required"`               // 用户名
	RealName string `v:"required"`               // 真实姓名
	Password string `v:"password2"`              // 密码
	Mobile   string ``                           // 手机号码
	Email    string ``                           // 邮箱
	Status   int    `v:"required|between:0,1"`   // 状态（1-开启，0-关闭）
	Avatar   string ``                           // 头像
	Roles    string `v:"required"`               // 角色集合
	GroupId  int    `v:"required"`               // 组织架构
}

type AdminUpdateInput struct {
	Id       int
	UserName string
	RealName string
	Password string
	Mobile   string
	Email    string
	Status   int
	Avatar   string
	Roles    string
	GroupId  int
}

type AdminDeleteReq struct {
	Id int `v:"required|min:0|integer"` // 管理员ID
}
