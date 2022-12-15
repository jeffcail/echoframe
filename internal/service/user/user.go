package serviceuser

import (
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
func UpdateUser() {

}

// DelUser
func DelUser() {

}

// UserList
func UserList() {

}
