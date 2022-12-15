package models

import (
	"time"
)

type UserInfo struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT"`
	UserId    int64     `json:"user_id" xorm:"pk autoincr BIGINT"`
	Phone     string    `json:"phone" xorm:"not null comment('手机号') VARCHAR(50)"`
	Email     string    `json:"email" xorm:"not null comment('邮箱') VARCHAR(50)"`
	CreatedAt time.Time `json:"created_at" xorm:"created not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Version   int64     `json:"version" xorm:"version not null comment('版本号') BIGINT"`
}

func (u *UserInfo) TableName() string {
	return "user_info"
}
