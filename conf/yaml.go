package conf

import (
	_yaml "github.com/echo-scaffolding/core/conf/yaml"
)

type CoreConfig struct {
	Debug    bool
	HTTPBind string
	Mysql    struct {
		DbDsn   string
		ShowSql bool
	}
	Redis struct {
		RedisAddr string
		Password  string
		RedisDb   int
	}
	LoggerPath string
}

var YConf *CoreConfig

//YamlConfig
func YamlConfig() {
	c := &CoreConfig{}
	_yaml.LoadCoreConfig(c)
	YConf = c
}
