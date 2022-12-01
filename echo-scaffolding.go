package main

import (
	"github.com/echo-scaffolding/router"

	"github.com/echo-scaffolding/boot"
)

func main() {
	boot.Boot()
	router.RunHttpServer()
}
