package serviceauth

import (
	"fmt"

	"github.com/echo-scaffolding/internal/out"

	"github.com/echo-scaffolding/pkg/jwt"

	"github.com/echo-scaffolding/pkg/uber"

	daoauth "github.com/echo-scaffolding/internal/dao/auth"
	"github.com/echo-scaffolding/internal/input"
	"go.uber.org/zap"
)

// AuthLogin
func AuthLogin(param *input.AuthLoginInput) (*out.AuthLoginOut, error) {
	u, err := daoauth.CheckUsernameAndPassword(param.Username, param.Password)
	if err != nil {
		uber.EchoScaLog.Error(fmt.Sprintf("账号和密码错误 %v", param.Username), zap.Error(err))
		return nil, err
	}
	cl := &jwt.JwtClaims{
		ID:       u.Id,
		Username: u.Username,
	}
	token, err := jwt.GenerateToken(cl)
	if err != nil {
		return nil, err
	}
	lr := &out.AuthLoginOut{
		Token:    token,
		Username: u.Username,
	}
	return lr, nil
}
