package _ini

import (
	"fmt"
	"log"

	"github.com/echoframe/conf"

	"github.com/go-ini/ini"
)

//LoadCoreConfig
func LoadCoreConfig(path string) *conf.CoreConfig {
	var err error

	cfg, err := ini.Load(path)
	if err != nil {
		log.Fatal("Fail to parse conf", err)
	}

	// db
	db, err := cfg.GetSection("mysql")
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to get section 'mysql': %v", err))
	}

	// redis
	rd, err := cfg.GetSection("redis")
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to get section 'redis': %v", err))
	}

	// l
	l, err := cfg.GetSection("Logger")
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to get section 'logger': %v", err))
	}

	// jwt
	j, err := cfg.GetSection("Jwt")
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to get section 'logger': %v", err))
	}

	conf.NewCoreConfig()
	conf.Config.Debug = cfg.Section("").Key("RUN_MODE").MustBool()
	conf.Config.HTTPBind = cfg.Section("").Key("HTTPBind").MustString("")
	conf.Config.Mysql.DbDsn = db.Key("DbDsn").MustString("")
	conf.Config.Mysql.ShowSQL = db.Key("ShowSql").MustBool()
	conf.Config.Redis.RedisAddr = rd.Key("RedisAddr").MustString("")
	conf.Config.Redis.Password = rd.Key("Password").MustString("")
	conf.Config.Redis.RedisDb = rd.Key("RedisDb").MustInt()
	conf.Config.Logger.Path = l.Key("Path").MustString("")
	conf.Config.Logger.MaxSize = l.Key("MaxSize").MustInt()
	conf.Config.Logger.MaxAge = l.Key("MaxAge").MustInt()
	conf.Config.Logger.Compress = l.Key("Compress").MustBool()
	conf.Config.Logger.LocalTime = l.Key("LocalTime").MustBool()
	conf.Config.Jwt.EXPIRE = j.Key("Expire").MustInt64()
	conf.Config.Jwt.SECRET = j.Key("SECRET").MustString("")
	conf.Config.LevelDBPath = cfg.Section("").Key("LevelDBPath").MustString("")
	conf.Config.MongoDB = cfg.Section("").Key("MongoDB").MustString("")
	conf.Config.EsUrl = cfg.Section("").Key("EsUrl").MustString("")
	return conf.Config
}
