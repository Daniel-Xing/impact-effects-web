package cache

import (
	"log"

	"github.com/go-redis/redis"
)

// cache 缓存
var RedisClient *redis.Client

// initRedis
func initRedis() *redis.Client {
	//redis
	redisURL := "172.24.0.3:6379"
	// 创建Redis连接
	var client *redis.Client
	maxTryTime := 3
	for i := 0; i < maxTryTime; i++ {
		client = redis.NewClient(&redis.Options{
			Addr:     redisURL,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		pong, err := client.Ping().Result()
		log.Println(pong, err)
		if err != nil {
			continue
		} else {
			log.Println("Redis Connect success!")
			break
		}
	}

	return client
}

// NewRedisConnection
func NewRedisConnection() *redis.Client {
	RedisClient = initRedis()
	return RedisClient
}

// Getcache 获取缓存
func GetCache() *redis.Client {
	if _, err := RedisClient.Ping().Result(); err != nil {
		RedisClient.Close()
		RedisClient = NewRedisConnection()
	}
	return RedisClient
}
