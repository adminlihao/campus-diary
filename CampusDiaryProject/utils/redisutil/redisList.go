package redisutil

import (
	"context"
	"time"
)

//RedisLPush  redis列表从左边插入
func RedisLPush(key string, value int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err := RDB.LPush(ctx, key, value).Err()
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}


//RedisLRange redis按照索引获取列表元素
func RedisLRange(key string,start int64,end int64) ([]string,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	results, err := RDB.LRange(ctx, key, start, end).Result()
	defer cancel()
	if err != nil {
		return nil,err
	}
	return results,nil
}


