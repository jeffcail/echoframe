package main

import (
	"fmt"

	confnacos "github.com/echo-scaffolding/conf/nacos"

	confini "github.com/echo-scaffolding/conf/ini"

	confyaml "github.com/echo-scaffolding/conf/yaml"

	"github.com/echo-scaffolding/router"

	"github.com/echo-scaffolding/boot"
)

func main() {
	boot.Boot()
	fmt.Println(confnacos.NConfig())
	fmt.Println(confyaml.YConf)
	fmt.Println(confini.Config())
	router.RunHttpServer()
}
