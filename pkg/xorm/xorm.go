package _xorm

import (
	"github.com/echo-scaffolding/common/global"
	"github.com/echo-scaffolding/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// SetMysql
func CreateMysql() (*xorm.Engine, error) {
	mysql, err := xorm.NewEngine("mysql", conf.Config.Mysql.DbDsn)
	if err != nil {
		return nil, err
	}
	mysql.ShowSQL(conf.Config.Mysql.ShowSQL)
	err = mysql.Ping()
	global.CheckErr(err)
	return mysql, nil
}
