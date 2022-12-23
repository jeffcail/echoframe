package daoauth

import (
	"errors"

	"github.com/echo-scaffolding/common/db"
	"github.com/echo-scaffolding/internal/v1/models"
	"github.com/echo-scaffolding/utils"
)

// CheckUsernameAndPassword
func CheckUsernameAndPassword(username, passwd string) (*models.User, error) {
	u := &models.User{}
	has, err := db.Mysql.Where("username = ? AND password = ?", username, utils.GeneratePasswd(passwd)).Get(u)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("账号密码错误")
	}
	return u, nil
}
