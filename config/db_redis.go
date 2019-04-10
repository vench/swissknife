package config

import (
	"github.com/go-redis/redis"
)

//
func GetNewClientRedis(host string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

//
func GetListClientRedis() (clients []*redis.Client) {
	size := len(config.Redis.List)
	if size == 0 {
		return
	}
	// init clients
	clients = make([]*redis.Client, size)
	for n, host := range config.Redis.List {
		clients[n] = GetNewClientRedis(host)
	}
	return
}