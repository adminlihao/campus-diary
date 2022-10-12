package redisutil

import (
	"CampusDiaryProject/setting"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	RDB *redis.Client
)

// InitRedis 初始化连接
func InitRedis(cfg *setting.RedisConfig) (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Username: cfg.UserName,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = RDB.Ping(ctx).Result()
	return err
}

func RedisClose() {
	RDB.Close()
}



//RedisDel   删除key
func RedisDel(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err := RDB.Del(ctx, key).Err()
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

//RedisExpire 设置过期时间
func RedisExpire(key string, expTime int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err := RDB.Expire(ctx, key, time.Duration(expTime)*time.Second).Err()
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}








