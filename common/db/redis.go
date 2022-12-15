package db

import "gopkg.in/redis.v5"

var Rc *redis.Client

// SetRedis
func SetRedis(_rc *redis.Client) {
	Rc = _rc
}
