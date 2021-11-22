// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// GroupDao is the manager for logic model data accessing and custom defined data operations functions management.
type GroupDao struct {
	Table   string       // Table is the underlying table name of the DAO.
	Group   string       // Group is the database configuration group name of current DAO.
	Columns GroupColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// GroupColumns defines and stores column names for table group.
type GroupColumns struct {
	Id        string // 组织架构ID
	Name      string // 组织架构名称
	Pid       string // 组织架构上级ID
	Link      string // 组织架构路径
	CreatedAt string // 创建时间
	CreatedId string // 添加人
	UpdatedAt string // 最后更新时间
	UpdatedId string // 最后更新人
	DeletedAt string // 删除时间
	DeletedId string // 删除人
}

//  groupColumns holds the columns for table group.
var groupColumns = GroupColumns{
	Id:        "id",
	Name:      "name",
	Pid:       "pid",
	Link:      "link",
	CreatedAt: "created_at",
	CreatedId: "created_id",
	UpdatedAt: "updated_at",
	UpdatedId: "updated_id",
	DeletedAt: "deleted_at",
	DeletedId: "deleted_id",
}

// NewGroupDao creates and returns a new DAO object for table data access.
func NewGroupDao() *GroupDao {
	return &GroupDao{
		Group:   "default",
		Table:   "group",
		Columns: groupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GroupDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
