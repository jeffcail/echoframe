package models

import (
	"time"
)

type User struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT"`
	Username  string    `json:"username" xorm:"not null comment('用户名') VARCHAR(50)"`
	Password  string    `json:"password" xorm:"not null comment('密码') VARCHAR(255)"`
	CreatedAt time.Time `json:"created_at" xorm:"created not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Version   int64     `json:"version" xorm:"version not null comment('版本号') BIGINT"`
}

func (u *User) TableName() string {
	return "user"
}
