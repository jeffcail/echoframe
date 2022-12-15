package db

import (
	"log"
	"sync"
	"time"

	"github.com/go-xorm/xorm"
)

var (
	instance *xorm.Engine
	one      sync.Once
	Mysql    *MysqlEngine
)

type MysqlEngine struct{}

// NewMysqlEngine
func NewMysqlEngine() *MysqlEngine {
	one.Do(func() {
		Mysql = &MysqlEngine{}
	})
	return Mysql
}

// SetMysql
func (m *MysqlEngine) SetMysql(e *xorm.Engine) {
	instance = e
	go func() {
		for {
			instance.Ping()
			time.Sleep(1 * time.Hour)
		}
	}()
}

// Transaction
func (m *MysqlEngine) Transaction(fs ...func(s *xorm.Session) error) error {
	session := instance.NewSession()
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
