package confyaml

import (
	"github.com/echo-scaffolding/conf"
	_yaml "github.com/echo-scaffolding/core/conf/yaml"
)

//YamlConfig
func YamlConfig(path string) {
	conf.NewCoreConfig()
	_yaml.LoadCoreConfig(path, conf.Config)
}
