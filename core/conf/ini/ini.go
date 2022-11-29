package _ini

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
)

type CoreConfig struct {
	RunMode    bool
	HTTPBind   string
	LoggerPath string
	Slat       string
	DbDsn      string
	ShowSql    bool
	RedisAddr  string
	Password   string
	RedisDb    int
}

//LoadCoreConfig
func LoadCoreConfig() *CoreConfig {
	var err error

	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal("Fail to parse conf", err)
	}

	// echo-scaffolding
	s, err := cfg.GetSection("echo-scaffolding")
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to get section 'echo-scaffolding': %v", err))
	}

	// db
	db, err := cfg.GetSection("mysql")
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to get section 'mysql': %v", err))
	}

	//redis
	rd, err := cfg.GetSection("redis")
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to get section 'redis': %v", err))
	}

	return &CoreConfig{
		RunMode:    cfg.Section("").Key("RUN_MODE").MustBool(),
		HTTPBind:   s.Key("HTTPBind").MustString(""),
		LoggerPath: s.Key("LoggerPath").MustString(""),
		Slat:       s.Key("Slat").MustString(""),
		DbDsn:      db.Key("DbDsn").MustString(""),
		ShowSql:    db.Key("ShowSql").MustBool(),
		RedisAddr:  rd.Key("RedisAddr").MustString(""),
		Password:   rd.Key("Password").MustString(""),
		RedisDb:    rd.Key("RedisDb").MustInt(),
	}
}
