// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"bieshu-oa/app/dao/internal"
)

// groupDao is the data access object for table group.
// You can define custom methods on it to extend its functionality as you wish.
type groupDao struct {
	*internal.GroupDao
}

var (
	// Group is globally public accessible object for table group operations.
	Group = groupDao{
		internal.NewGroupDao(),
	}
)

// Fill with you ideas below.
