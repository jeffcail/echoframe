package boot

import (
	"github.com/echoframe/common/db"
	"github.com/echoframe/conf"
	_redis "github.com/echoframe/pkg/redis"
)

// InstanceRedis
func InstanceRedis() {
	addr := conf.Config.Redis.RedisAddr
	passwd := conf.Config.Redis.Password
	database := conf.Config.Redis.RedisDb
	client := _redis.CreateRedis(addr, passwd, database)
	db.SetRedis(client)
}
