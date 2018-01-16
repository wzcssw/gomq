package goqueue

import (
	"strconv"

	"github.com/go-redis/redis"
)

type Goqueue struct {
	// redis 配置
	RedisConf RedisConfig

	// redis 连接   form: github.com/go-redis/redis
	RedisClient *redis.Client
}

// 获取队列长度
func (queue *Goqueue) GetQueueLength(key string) (int64, error) {
	intCmd := queue.RedisClient.LLen(key)
	return intCmd.Val(), intCmd.Err()
}

// 获取列表指定范围内的元素
func (queue *Goqueue) GetQueueRange(key string, start int64, stop int64) ([]string, error) {
	stringSliceCmd := queue.RedisClient.LRange(key, start, stop)
	return stringSliceCmd.Val(), stringSliceCmd.Err()
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func (rc *RedisConfig) Addr() string {
	port := strconv.Itoa(rc.Port)

	return rc.Host + ":" + port
}

// 获得Gomq实例
func NewInstance(rc RedisConfig) *Goqueue {
	var gomq Goqueue
	// 得到Redis连接
	redisClient := newRedisClient(rc)
	gomq.RedisClient = redisClient

	return &gomq
}

func newRedisClient(rc RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     rc.Addr(),
		Password: rc.Password, // no password set
		DB:       rc.DB,       // use default DB
	})
}
