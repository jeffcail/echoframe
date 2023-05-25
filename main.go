package main

import (
	"github.com/echoframe/boot"
	"github.com/echoframe/router"
)

func main() {
	boot.Boot()
	router.RunHttpServer()
}
