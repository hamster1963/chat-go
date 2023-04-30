// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"push/internal/dao/internal"
)

// internalPushServiceDao is internal type for wrapping internal DAO implements.
type internalPushServiceDao = *internal.PushServiceDao

// pushServiceDao is the data access object for table push_service.
// You can define custom methods on it to extend its functionality as you wish.
type pushServiceDao struct {
	internalPushServiceDao
}

var (
	// PushService is globally public accessible object for table push_service operations.
	PushService = pushServiceDao{
		internal.NewPushServiceDao(),
	}
)

// Fill with you ideas below.
