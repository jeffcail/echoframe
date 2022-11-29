package main

import (
	"fmt"

	confnacos "github.com/echo-scaffolding/conf/nacos"

	confyaml "github.com/echo-scaffolding/conf/yaml"

	confini "github.com/echo-scaffolding/conf/ini"

	"github.com/echo-scaffolding/boot"
)

func main() {
	boot.Bootstrap()
	fmt.Println(confini.Config())
	fmt.Println(confyaml.YConf)
	fmt.Println(confnacos.NConfig())
}
