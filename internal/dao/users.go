// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"Capys/internal/dao/internal"
)

// internalUsersDao is an internal type for wrapping the internal DAO implementation.
type internalUsersDao = *internal.UsersDao

// usersDao is the data access object for the table users.
// You can define custom methods on it to extend its functionality as needed.
type usersDao struct {
	internalUsersDao
}

var (
	// Users is a globally accessible object for table users operations.
	Users = usersDao{
		internal.NewUsersDao(),
	}
)

// Add your custom methods and functionality below.
