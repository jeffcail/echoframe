package main

import (
	"log"

	"github.com/echo-scaffolding/boot"
	"github.com/echo-scaffolding/conf"
)

func main() {
	boot.Bootstrap()
	log.Println(conf.YConf.Mysql.DbDsn)
}
