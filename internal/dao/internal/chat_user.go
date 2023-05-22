// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChatUserDao is the data access object for table chat_user.
type ChatUserDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns ChatUserColumns // columns contains all the column names of Table for convenient usage.
}

// ChatUserColumns defines and stores column names for table chat_user.
type ChatUserColumns struct {
	Id           string // 用户ID
	Username     string // 用户名
	LastChatTime string // 上次聊天时间
	UserStatus   string // 用户状态
}

// chatUserColumns holds the columns for table chat_user.
var chatUserColumns = ChatUserColumns{
	Id:           "id",
	Username:     "username",
	LastChatTime: "last_chat_time",
	UserStatus:   "user_status",
}

// NewChatUserDao creates and returns a new DAO object for table data access.
func NewChatUserDao() *ChatUserDao {
	return &ChatUserDao{
		group:   "default",
		table:   "chat_user",
		columns: chatUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChatUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChatUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChatUserDao) Columns() ChatUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChatUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChatUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChatUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
