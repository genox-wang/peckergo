package datacache

import (
	"time"

	// "github.com/patrickmn/go-cache"
	"peckergo/api/config"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// var (
// 	_cache *cache.Cache
// )

// func init() {
// 	_cache = cache.New(cache.DefaultExpiration, time.Minute*1)
// }
var (
	_cache *redis.Client
)

func init() {
	_cache = redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.Addr"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := _cache.Ping().Result()
	if err != nil {
		log.Panicf("datacahe init : %s\n", err.Error())
	}
	// InitchannelUUIDCache()
	// InitsourceIDCache()
	// InitAdStrategyByIDCache()
	// _cache = cache.New(cache.DefaultExpiration, time.Minute*1)
}

// Set 缓存设置键值对
func Set(k string, v interface{}, expire time.Duration) {
	_cache.Set(k, v, expire)
}

// Get 获取缓存
func Get(k string) (string, bool) {
	// return _cache.Get(k)
	r := _cache.Get(k)
	return r.Val(), r.Err() == nil
}

// SetNX 缓存添加键值对，存在返回 false
func SetNX(k string, v interface{}, expire time.Duration) bool {
	r := _cache.SetNX(k, v, expire)
	success, _ := r.Result()
	// fmt.Printf("%s_%+v\n", k, success)
	return success
}

// Incr 缓存 自增
func Incr(k string, n int64) (int64, error) {
	r := _cache.IncrBy(k, n)
	return r.Val(), r.Err()
}

// Decr 缓存 自减
func Decr(k string, n int64) (int64, error) {
	r := _cache.DecrBy(k, n)
	return r.Val(), r.Err()
}

// Delete 删除键
func Delete(k string) {
	_cache.Del(k)
}

// HSet hset
func HSet(k, f string, v interface{}) {
	_cache.HSet(k, f, v)
}

// HSetNX 缓存添加键值对，存在返回 false
func HSetNX(k, f string, v interface{}) bool {
	r := _cache.HSetNX(k, f, v)
	success, _ := r.Result()
	return success
}

// HIncr HIncr
func HIncr(k, f string, n int64) (int64, error) {
	r := _cache.HIncrBy(k, f, n)
	return r.Result()
}

// Expire 为兼设置过期时间
func Expire(k string, expire time.Duration) {
	_cache.Expire(k, expire)
}

// TTL 获取过期时间
func TTL(k string) time.Duration {
	r := _cache.TTL(k)
	return r.Val()
}

// HExists hexists
func HExists(k, f string) bool {
	r, _ := _cache.HExists(k, f).Result()
	return r
}

// HGet hget
func HGet(k, f string) (string, error) {
	r := _cache.HGet(k, f)
	return r.Val(), r.Err()
}

// HDel hdel
func HDel(k, f string) {
	_cache.HDel(k, f)
}
