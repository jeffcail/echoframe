package boot

import (
	confini "github.com/echoframe/conf/ini"
)

//InitIni
func InitIni(path string) {
	confini.ParseIniConfig(path)
}
