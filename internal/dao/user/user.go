package daouser

import (
	"errors"

	"github.com/echo-scaffolding/common/db"
	"github.com/echo-scaffolding/internal/input"
	"github.com/echo-scaffolding/internal/v1/models"
	"github.com/echo-scaffolding/utils"
	"github.com/go-xorm/xorm"
)

// CreateUser
func CreateUser(param *input.CreateUserInput) error {
	err := db.Transaction(func(s *xorm.Session) error {
		u := &models.User{
			Username: param.Username,
			Password: utils.GeneratePasswd(param.Password),
		}
		_, err := s.Insert(u)
		if err != nil {
			return err
		}

		ui := &models.UserInfo{
			UserId: u.Id,
			Phone:  param.Phone,
			Email:  param.Email,
		}
		_, err = s.Insert(ui)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

// UserDetail
func UserDetail(id int64) (*models.User, error) {
	u := &models.User{}
	has, err := db.Mysql.ID(id).Get(u)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户不存在")
	}
	return u, nil
}

// UpdateUser
func UpdateUser() {

}

// DelUser
func DelUser() {

}

// UserList
func UserList() {

}
