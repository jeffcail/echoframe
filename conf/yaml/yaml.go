package confyaml

import (
	"github.com/echoframe/conf"
	_yaml "github.com/echoframe/core/conf/yaml"
)

//YamlConfig
func YamlConfig(path string) {
	conf.NewCoreConfig()
	_yaml.LoadCoreConfig(path, conf.Config)
}
