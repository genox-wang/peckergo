package mq

import (
	"peckergo/api/config"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
)

const (
	redisMQPrefix = "mq"
)

type redisMQ struct {
	client *redis.Client
}

// Init 初始化
func (r *redisMQ) Init() {
	r.client = redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.Addr"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := r.client.Ping().Result()
	if err != nil {
		log.Panicf("redisMQ init : %s\n", err.Error())
	}
}

// Produce 生产
func (r *redisMQ) Produce(topic string, data string) {
	r.client.RPush(fmt.Sprintf("%s_%s", redisMQPrefix, topic), data)
}

// Consume 消费
func (r *redisMQ) Consume(topic string, maxLine int) []string {
	key := fmt.Sprintf("%s_%s", redisMQPrefix, topic)
	ss := r.client.LRange(key, 0, int64(maxLine))
	s := ss.Val()
	r.client.LTrim(key, int64(len(s)), -1)
	return s
}
