package confnacos

import (
	"flag"

	"github.com/echo-scaffolding/core/conf/driver"

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

type NacosConfig struct {
	Debug      bool
	HTTPBind   string
	Mysql      driver.MysqlConfig
	Redis      driver.RedisConfig
	LoggerPath string
}

var config NacosConfig

func NConfig() NacosConfig {
	return config
}

//InitNacos
func InitNacos() {
	_nacos.LoadCoreConfig(*ip, *port, *cfg, *group, &config)
}
