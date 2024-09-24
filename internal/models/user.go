package models

import (
	"time"
)

// User represents a row in the 'mining_user' table.
type User struct {
	Id       int64     `xorm:"'id' pk autoincr"`
	Username string    `xorm:"username"`
	Created  time.Time `xorm:"'created'"` // 矿机用户创建时间
	Updated  time.Time `xorm:"'updated'"` // 矿机用户修改时间
	Deleted  time.Time `xorm:"'deleted'"` // 矿机用户删除时间
	Version  int64     `xorm:"'version'"` // 版本号
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}
