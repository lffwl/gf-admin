// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"bieshu-oa/app/dao/internal"
)

// roleMenuDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type roleMenuDao struct {
	*internal.RoleMenuDao
}

var (
	// RoleMenu is globally public accessible object for table role_menu operations.
	RoleMenu = roleMenuDao{
		internal.NewRoleMenuDao(),
	}
)

// Fill with you ideas below.
