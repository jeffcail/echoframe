package conf

import _ini "github.com/echo-scaffolding/core/conf/ini"

var config *_ini.CoreConfig

//Config
func Config() *_ini.CoreConfig {
	return config
}

//ParseIniConfig
func ParseIniConfig() {
	config = _ini.LoadCoreConfig()
}
