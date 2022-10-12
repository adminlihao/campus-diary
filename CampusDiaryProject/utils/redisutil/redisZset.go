package redisutil

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

//RedisZAdd  redis添加zset数据
func RedisZAdd(key string, score float64, member string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err := RDB.ZAdd(ctx, key, &redis.Z{
		Score:  score,
		Member: member,
	}).Err()
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

//RedisZRem  redis删除zset的值
func RedisZRem(key string, member string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err := RDB.ZRem(ctx, key, member).Err()
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

//RedisZCard   redis获取zset元素个数
func RedisZCard(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := RDB.ZCard(ctx, key).Result()
	defer cancel()
	if err != nil {
		return -1, err
	}
	return result, err
}

//RedisZRevRange  redis获取根据score倒序排序后的数据段
func RedisZRevRange(key string, start int64, len int64) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	results, err := RDB.ZRevRange(ctx, key, start, len).Result()
	defer cancel()
	if err != nil {
		return nil, err
	}
	return results, nil
}

//RedisZRevRangeWithScores  redis获取根据score倒序排序后的数据段包含score
func RedisZRevRangeWithScores(key string, start int64, len int64) ([]redis.Z, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	results, err := RDB.ZRevRangeWithScores(ctx, key, start, len).Result()
	defer cancel()
	if err != nil {
		return nil, err
	}
	return results, nil
}


//RedisZRank   redis根据ZRank实现检查元素是否存在
func RedisZRank(key string, member string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := RDB.ZRank(ctx, key, member).Result()
	defer cancel()
	log.Println(result)
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return -1, err
	}
	if result >= 0 {
		return 1, nil
	}
	return -1, nil
}


