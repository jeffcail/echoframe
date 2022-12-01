package boot

import (
	confini "github.com/echo-scaffolding/conf/ini"
)

//InitIni
func InitIni(path string) {
	confini.ParseIniConfig(path)
}
