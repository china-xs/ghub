// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser2role = "a_user2role"

// User2role mapped from table <a_user2role>
type User2role struct {
	ID        int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`      // ID
	AccountID int32 `gorm:"column:account_id;primaryKey" json:"account_id"`         // account.id
	RoleID    int32 `gorm:"column:role_id;primaryKey" json:"role_id"`               // role.id
	OperateID int32 `gorm:"column:operate_id;not null;default:0" json:"operate_id"` // 操作人ID
}

// TableName User2role's table name
func (*User2role) TableName() string {
	return TableNameUser2role
}
