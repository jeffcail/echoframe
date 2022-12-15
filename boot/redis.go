package boot

import (
	"github.com/echo-scaffolding/common/db"
	"github.com/echo-scaffolding/conf"
	_redis "github.com/echo-scaffolding/pkg/redis"
)

// InstanceRedis
func InstanceRedis() {
	addr := conf.Config.Redis.RedisAddr
	passwd := conf.Config.Redis.Password
	database := conf.Config.Redis.RedisDb
	client := _redis.CreateRedis(addr, passwd, database)
	db.SetRedis(client)
}
