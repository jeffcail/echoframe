package models

import (
	"time"
)

type User struct {
	Id        int64     `xorm:"pk autoincr BIGINT"`
	Username  string    `xorm:"not null comment('用户名') VARCHAR(50)"`
	Password  string    `xorm:"not null comment('密码') VARCHAR(50)"`
	CreatedAt time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Version   int64     `xorm:"not null comment('版本号') BIGINT"`
}

func (u *User) TableName() string {
	return "user"
}
