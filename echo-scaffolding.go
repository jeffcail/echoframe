package main

import (
	"github.com/echo-scaffolding/boot"
	"github.com/echo-scaffolding/conf"
)

func main() {
	boot.Bootstrap()
	println(conf.Config().DbDsn)
}
