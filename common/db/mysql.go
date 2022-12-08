package db

import (
	"log"
	"time"

	"github.com/go-xorm/xorm"
)

var Mysql *xorm.Engine

type MysqlEngine struct{}

// NewMysqlEngine
func NewMysqlEngine() *MysqlEngine {
	return &MysqlEngine{}
}

// SetMysql
func (m *MysqlEngine) SetMysql(e *xorm.Engine) {
	Mysql = e
	go func() {
		for {
			Mysql.Ping()
			time.Sleep(1 * time.Hour)
		}
	}()
}

// Transaction
func (m *MysqlEngine) Transaction(fs ...func(s *xorm.Session) error) error {
	session := Mysql.NewSession()
	session.Begin()
	for _, f := range fs {
		err := f(session)
		if err != nil {
			log.Println(err)
			session.Rollback()
			session.Clone()
			return err
		}
	}
	session.Commit()
	session.Clone()
	return nil
}

// Close
func (m *MysqlEngine) Close() {
	Mysql.Close()
}
