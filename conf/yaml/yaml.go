package confyaml

import (
	"github.com/echo-scaffolding/core/conf/driver"
	_yaml "github.com/echo-scaffolding/core/conf/yaml"
)

type CoreConfig struct {
	Debug      bool
	HTTPBind   string
	Mysql      driver.MysqlConfig
	Redis      driver.RedisConfig
	LoggerPath string
}

var YConf *CoreConfig

//YamlConfig
func YamlConfig() {
	c := &CoreConfig{}
	_yaml.LoadCoreConfig(c)
	YConf = c
}
