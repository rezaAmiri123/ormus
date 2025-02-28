package rediswritekey

import "github.com/rezaAmiri123/ormus/adapter/redis"

type DB struct {
	adapter redis.Adapter
}

// New is Constructor redis DB.
func New(adapter redis.Adapter) DB {
	return DB{adapter: adapter}
}
