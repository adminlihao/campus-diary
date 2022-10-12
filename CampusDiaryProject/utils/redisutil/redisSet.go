package redisutil

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

//RedisSAdd redis添加set集合元素
func RedisSAdd(key string, member string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err := RDB.SAdd(ctx, key, member).Err()
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

//RedisSRem redis删除set集合元素
func RedisSRem(key string, member string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err := RDB.SRem(ctx, key, member).Err()
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

//RedisSMembers redis获取set集合所有元素
func RedisSMembers(key string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	results, err := RDB.SMembers(ctx, key).Result()
	defer cancel()
	if err != nil {
		return nil, err
	}
	return results, nil
}

//RedisSIsMember   redis判断元素是否在set中
func RedisSIsMember(key string,member string) (bool,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	exists, err := RDB.SIsMember(ctx, key, member).Result()
	defer cancel()
	if err != nil {
		return false,err
	}
	return exists, nil
}

//RedisSRandMember redis随机取出set中1个值
func RedisSRandMember (key string) (string,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := RDB.SRandMember(ctx, key).Result()
	defer cancel()
	if err ==redis.Nil{
		return "",nil
	}
	if err != nil {
		return "",err
	}
	return result,nil
}

//RedisSRandMemberN redis随机取出set中n个值
func RedisSRandMemberN(key string,count int64) ([]string,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := RDB.SRandMemberN(ctx, key, count).Result()
	defer cancel()
	if err ==redis.Nil{
		return nil,nil
	}
	if err != nil {
		return nil,err
	}
	return result,nil
}
