package boot2

import (
	"github.com/echoframe/common/db"
	"github.com/echoframe/common/global"
	_xorm "github.com/echoframe/pkg/xorm"
	"github.com/go-xorm/xorm"
)

// InitMysql
func InitMysql() {
	engine, err := _xorm.CreateMysql()
	global.CheckErr(err)
	engine.SetLogger(xorm.NewSimpleLogger(_xorm.XLog))
	db.SetMysql(engine)
}
