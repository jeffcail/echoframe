package confnacos

import (
	"flag"

	"github.com/echo-scaffolding/conf"

	_nacos "github.com/echo-scaffolding/core/conf/nacos"
)

var (
	ip    = flag.String("ip", "ip", "The nacos of ip address")
	port  = flag.Int("p", 0, "The nacos of port")
	cfg   = flag.String("c", "default", "The nacos of Data ID")
	group = flag.String("g", "default", "The nacos of Group")
)

func init() {
	flag.Parse()
}

//InitNacos
func InitNacos() {
	conf.NewCoreConfig()
	_nacos.LoadCoreConfig(*ip, *port, *cfg, *group, conf.Config)
}
