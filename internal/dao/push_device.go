// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"push/internal/dao/internal"
)

// internalPushDeviceDao is internal type for wrapping internal DAO implements.
type internalPushDeviceDao = *internal.PushDeviceDao

// pushDeviceDao is the data access object for table push_device.
// You can define custom methods on it to extend its functionality as you wish.
type pushDeviceDao struct {
	internalPushDeviceDao
}

var (
	// PushDevice is globally public accessible object for table push_device operations.
	PushDevice = pushDeviceDao{
		internal.NewPushDeviceDao(),
	}
)

// Fill with you ideas below.