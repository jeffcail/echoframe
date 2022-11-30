package boot

import (
	"os"
)

func Bootstrap() {
	InitIni()
	InitYaml()
	cmd()
	InitLogger()
}

func cmd() {
	if len(os.Args) <= 1 {
		return
	}
	if os.Args[1] == "-ip" && os.Args[3] == "-p" && os.Args[5] == "-c" && os.Args[7] == "-g" {
		InitNacos()
	}
}
