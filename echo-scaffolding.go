package main

import (
	"github.com/echo-scaffolding/boot"
	"github.com/echo-scaffolding/router"
)

func main() {
	boot.Boot()

	router.RunHttpServer()
}
