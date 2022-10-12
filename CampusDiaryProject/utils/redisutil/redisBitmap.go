package redisutil

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

//RedisSetBit    redis写入bitmap数据
func RedisSetBit(key string, offset int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err := RDB.SetBit(ctx, key, offset, 1).Err()
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

//RedisGetBit    redis获取bitmap数据
func RedisGetBit(key string,offset int64) (int64,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := RDB.GetBit(ctx, key, offset).Result()
	defer cancel()
	if err==redis.Nil{
		return -1,nil
	}
	if err != nil {
		return -1,err
	}
	return result,nil
}
