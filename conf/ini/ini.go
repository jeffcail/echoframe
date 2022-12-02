package confini

import (
	_ini "github.com/echo-scaffolding/core/conf/ini"
)

//ParseIniConfig
func ParseIniConfig(path string) {
	_ini.LoadCoreConfig(path)
}
