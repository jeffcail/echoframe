package confyaml

import (
	"github.com/echo-scaffolding/core/conf/driver"
	_yaml "github.com/echo-scaffolding/core/conf/yaml"
)

type CoreConfig struct {
	Debug    bool
	HTTPBind string
	Mysql    driver.MysqlConfig
	Redis    driver.RedisConfig
	Logger   struct {
		Path      string
		MaxSize   int
		MaxAge    int
		Compress  bool
		LocalTime bool
	}
}

var YConf *CoreConfig

//YamlConfig
func YamlConfig() {
	c := &CoreConfig{}
	_yaml.LoadCoreConfig(c)
	YConf = c
}
