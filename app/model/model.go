// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

import (
	"github.com/gogf/gf/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	Id        int         `orm:"id,primary"       json:"id"`         // 管理员ID
	UserName  string      `orm:"user_name,unique" json:"user_name"`  // 用户名
	RealName  string      `orm:"real_name"        json:"real_name"`  // 真实姓名
	Password  string      `orm:"password"         json:"password"`   // 密码
	Mobile    string      `orm:"mobile"           json:"mobile"`     // 手机号码
	Email     string      `orm:"email"            json:"email"`      // 邮箱
	Status    int         `orm:"status"           json:"status"`     // 状态（1-开启，0-关闭）
	Avatar    string      `orm:"avatar"           json:"avatar"`     // 头像
	GroupId   int         `orm:"group_id"         json:"group_id"`   // 组织架构
	CreatedAt *gtime.Time `orm:"created_at"       json:"created_at"` // 创建时间
	CreatedId int         `orm:"created_id"       json:"created_id"` // 添加人
	UpdatedAt *gtime.Time `orm:"updated_at"       json:"updated_at"` // 最后更新时间
	UpdatedId int         `orm:"updated_id"       json:"updated_id"` // 最后更新人
	DeletedAt *gtime.Time `orm:"deleted_at"       json:"deleted_at"` // 删除时间
	DeletedId int         `orm:"deleted_id"       json:"deleted_id"` // 删除人
}

// AdminRole is the golang structure for table admin_role.
type AdminRole struct {
	AdminId int `orm:"admin_id,primary" json:"admin_id"` // 管理员ID
	RoleId  int `orm:"role_id,primary"  json:"role_id"`  // 角色ID
}

// Group is the golang structure for table group.
type Group struct {
	Id        int         `orm:"id,primary" json:"id"`         // 组织架构ID
	Name      string      `orm:"name"       json:"name"`       // 组织架构名称
	Pid       int         `orm:"pid"        json:"pid"`        // 组织架构上级ID
	Link      string      `orm:"link"       json:"link"`       // 组织架构路径
	CreatedAt *gtime.Time `orm:"created_at" json:"created_at"` // 创建时间
	CreatedId int         `orm:"created_id" json:"created_id"` // 添加人
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updated_at"` // 最后更新时间
	UpdatedId int         `orm:"updated_id" json:"updated_id"` // 最后更新人
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deleted_at"` // 删除时间
	DeletedId int         `orm:"deleted_id" json:"deleted_id"` // 删除人
}

// Menu is the golang structure for table menu.
type Menu struct {
	Id        int         `orm:"id,primary" json:"id"`         // 菜单ID
	Name      string      `orm:"name"       json:"name"`       // 菜单名称
	Router    string      `orm:"router"     json:"router"`     // 请求地址
	Method    int         `orm:"method"     json:"method"`     // 请求类型(1 = "Get",2 = "Post",3 = "Put",4 = "DELETE")
	Key       string      `orm:"key"        json:"key"`        // 标识
	Pid       int         `orm:"pid"        json:"pid"`        // 菜单上级ID
	Type      int         `orm:"type"       json:"type"`       // 菜单类型（1= "目录"2 = “菜单”，3 = “按钮”，4 = “隐藏”）
	Link      string      `orm:"link"       json:"link"`       // 菜单路径
	CreatedAt *gtime.Time `orm:"created_at" json:"created_at"` // 创建时间
	CreatedId int         `orm:"created_id" json:"created_id"` // 添加人
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updated_at"` // 最后更新时间
	UpdatedId int         `orm:"updated_id" json:"updated_id"` // 最后更新人
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deleted_at"` // 删除时间
	DeletedId int         `orm:"deleted_id" json:"deleted_id"` // 删除人
}

// Role is the golang structure for table role.
type Role struct {
	Id        int         `orm:"id,primary" json:"id"`         // 角色ID
	Name      string      `orm:"name"       json:"name"`       // 角色名称
	Dp        int         `orm:"dp"         json:"dp"`         // 数据权限（0-仅自己，1-所在组织架构和下级组织架构，2-所有人）
	CreatedAt *gtime.Time `orm:"created_at" json:"created_at"` // 创建时间
	CreatedId int         `orm:"created_id" json:"created_id"` // 添加人
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updated_at"` // 最后更新时间
	UpdatedId int         `orm:"updated_id" json:"updated_id"` // 最后更新人
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deleted_at"` // 删除时间
	DeletedId int         `orm:"deleted_id" json:"deleted_id"` // 删除人
}

// RoleMenu is the golang structure for table role_menu.
type RoleMenu struct {
	RoleId int `orm:"role_id,primary" json:"role_id"` // 角色ID
	MenuId int `orm:"menu_id,primary" json:"menu_id"` // 菜单ID
}
