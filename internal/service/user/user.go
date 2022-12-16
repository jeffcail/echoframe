package serviceuser

import (
	"fmt"

	"github.com/echo-scaffolding/internal/filter"

	_estime "github.com/echo-scaffolding/common/estime"
	daouser "github.com/echo-scaffolding/internal/dao/user"
	"github.com/echo-scaffolding/internal/input"
	"github.com/echo-scaffolding/internal/out"
	"github.com/echo-scaffolding/pkg/uber"
	"go.uber.org/zap"
)

// CreateUser
func CreateUser(param *input.CreateUserInput) error {
	if err := daouser.CreateUser(param); err != nil {
		uber.EchoScaLog.Error("创建user账号失败")
		return err
	}
	return nil
}

// UserDetail
func UserDetail(id int64) (*out.UserDetailOut, error) {
	u, err := daouser.UserDetail(id)
	if err != nil {
		uber.EchoScaLog.Error("用户不存在", zap.Int64("id", id), zap.Error(err))
		return nil, err
	}

	o := &out.UserDetailOut{
		ID:        u.Id,
		Username:  u.Username,
		CreatedAt: _estime.FormatTime(u.CreatedAt),
		UpdatedAt: _estime.FormatTime(u.UpdatedAt),
	}

	return o, nil
}

// UpdateUser
func UpdateUser(param *input.UpdateUserInput) (err error) {
	if err = daouser.UpdateUser(param); err != nil {
		uber.EchoScaLog.Error(fmt.Sprintf("id【%d】修改信息失败", param.ID), zap.Error(err))
		return
	}
	return
}

// DelUser
func DelUser(id int64) (err error) {
	if err = daouser.DelUser(id); err != nil {
		uber.EchoScaLog.Error(fmt.Sprintf("id【%d】删除失败", id), zap.Error(err))
		return
	}
	return
}

// UserList
func UserList(param *input.UserListInput) (int64, []*out.UserListOut, error) {
	count, users, err := daouser.UserList(param, filter.UserListFilter(param.Username))
	if err != nil {
		uber.EchoScaLog.Error("获取用户信息列表失败", zap.Error(err))
		return 0, nil, err
	}
	us := make([]*out.UserListOut, 0)
	for _, v := range users {
		u := &out.UserListOut{
			ID:        v.Id,
			Username:  v.Username,
			CreatedAt: _estime.FormatTime(v.CreatedAt),
			UpdatedAt: _estime.FormatTime(v.UpdatedAt),
		}

		us = append(us, u)
	}
	return count, us, nil
}
