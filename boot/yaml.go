package boot

import confyaml "github.com/echoframe/conf/yaml"

//InitYaml
func InitYaml(path string) {
	confyaml.YamlConfig(path)
}
