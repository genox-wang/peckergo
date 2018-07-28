package datacache

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Set 设置缓存键值对
func Set(key string, value interface{}, expTime time.Duration) {
	client.Set(key, value, expTime)
}

// Get 通过键获取缓存数据
func Get(key string) (string, error) {
	return client.Get(key).Result()
}

// Del 通过键删除缓存
func Del(key string) error {
	_, err := client.Del(key).Result()
	return err
}
