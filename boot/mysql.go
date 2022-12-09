package boot

import (
	"github.com/echo-scaffolding/common/db"
	"github.com/echo-scaffolding/common/global"
	_xorm "github.com/echo-scaffolding/pkg/xorm"
	"github.com/go-xorm/xorm"
)

// InitMysql
func InitMysql() {
	engine, err := _xorm.CreateMysql()
	global.CheckErr(err)
	engine.SetLogger(xorm.NewSimpleLogger(_xorm.XLog))
	db.Mysql.SetMysql(engine)
}
