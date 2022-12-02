package main

import (
	"fmt"

	"github.com/echo-scaffolding/conf"
	"github.com/echo-scaffolding/router"

	"github.com/echo-scaffolding/boot"
)

func main() {
	boot.Boot()
	fmt.Println(conf.Config)
	router.RunHttpServer()
}
