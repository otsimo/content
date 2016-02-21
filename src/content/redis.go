package content

import (
	redis "gopkg.in/redis.v3"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(config *Config) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	return &RedisClient{
		client: client,
	}
}
