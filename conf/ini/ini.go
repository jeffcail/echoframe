package confini

import (
	_ini "github.com/echoframe/core/conf/ini"
)

//ParseIniConfig
func ParseIniConfig(path string) {
	_ini.LoadCoreConfig(path)
}
