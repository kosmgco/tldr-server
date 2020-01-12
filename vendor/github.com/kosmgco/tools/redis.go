package tools

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	PoolSize int    `json:"poolSize"`
	db       *redis.Client
}

func (r *Redis) Get() *redis.Client {
	if r.db == nil {
		c := redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     fmt.Sprintf("%s:%s", r.Host, r.Port),
			PoolSize: r.PoolSize,
		})
		r.db = c
		return c
	}
	return r.db
}

func (r *Redis) Close() {
	if r != nil && r.db != nil {
		r.db = nil
	}
}
