package daouser

import (
	"errors"

	"github.com/echoframe/common/db"
	"github.com/echoframe/internal/input"
	"github.com/echoframe/internal/v1/models"
	"github.com/echoframe/utils"
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
	has, err := db.Mysql.ID(id).ForUpdate().Get(u)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户不存在")
	}
	return u, nil
}

// UpdateUser
func UpdateUser(param *input.UpdateUserInput) error {
	u, err := UserDetail(param.ID)
	if err != nil {
		return err
	}

	err = db.Transaction(func(s *xorm.Session) error {
		u.Username = param.Username
		_, err = s.ID(u.Id).ForUpdate().Update(u)
		if err != nil {
			return err
		}

		ui := &models.UserInfo{}
		_, err = s.Where("user_id = ?", u.Id).ForUpdate().Get(ui)
		if err != nil {
			return err
		}
		ui.Phone = param.Phone
		ui.Email = param.Email
		_, err = s.Where("user_id = ?", u.Id).ForUpdate().Update(ui)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// DelUser
func DelUser(id int64) (err error) {
	u, err := UserDetail(id)
	if err != nil {
		return
	}
	err = db.Transaction(func(s *xorm.Session) error {
		_, err = s.ID(u.Id).ForUpdate().Delete(u)
		if err != nil {
			return err
		}
		ui := &models.UserInfo{}
		_, err = s.Where("user_id = ?", u.Id).ForUpdate().Get(ui)
		if err != nil {
			return err
		}

		_, err = s.Where("user_id = ?", u.Id).ForUpdate().Delete(ui)
		if err != nil {
			return err
		}

		return nil
	})
	return
}

// UserList
func UserList(param *input.UserListInput, filter map[string]interface{}) (int64, []*models.User, error) {
	data := make([]*models.User, 0)
	var query *xorm.Session
	var query2 *xorm.Session

	if param.Page != -1 {
		query = db.Mysql.Limit(param.PageSize, (param.Page-1)*param.PageSize)
	}
	query2 = db.Mysql.NewSession()
	for k, v := range filter {
		query.Where(k+"=?", v)
		query2.Where(k+"=?", v)
	}
	count, _ := query2.Count(&models.User{})
	query.Desc("id")
	err := query.Find(&data)
	return count, data, err
}
