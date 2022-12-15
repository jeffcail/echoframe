package _redis

import (
	"github.com/echo-scaffolding/common/global"
	"gopkg.in/redis.v5"
)

// CreateRedis
func CreateRedis(addr, passwd string, db int) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	_, err := redisClient.Ping().Result()
	global.CheckErr(err)
	return redisClient
}
