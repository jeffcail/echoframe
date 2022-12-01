package boot

import confyaml "github.com/echo-scaffolding/conf/yaml"

//InitYaml
func InitYaml(path string) {
	confyaml.YamlConfig(path)
}
