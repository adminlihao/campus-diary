package service

import (
	"CampusDiaryProject/models"
	"CampusDiaryProject/utils/redisutil"
	"strconv"
)

//GetRelationInfoByUserId  获取好友信息
func GetRelationInfoByUserId(key string, field string) (*models.RelationInfo, error) {
	val, err := redisutil.RedisHGet(models.RELATION+key, field)
	if err != nil {
		return nil, err
	}
	if val == "" {
		val = "0"
	}
	status, _ := strconv.Atoi(val)
	userId, _ := strconv.Atoi(field)
	userInfo, _ := GetAUserInfoByUserId(userId)
	friendInfo := &models.RelationInfo{
		Status:   status,
		UserInfo: userInfo,
	}
	return friendInfo, nil
}

//GetRelationInfoStatus   获取和好友之间状态
func GetRelationInfoStatus(key string, field string) (int, error) {
	val, err := redisutil.RedisHGet(models.RELATION+key, field)
	if err != nil {
		return -1, err
	}
	if val == "" {
		val = "0"
	}
	status, _ := strconv.Atoi(val)
	return status, nil
}

//ChangeRelationInfoStatus  改变和好友之间的状态状态
func ChangeRelationInfoStatus(key string, field string, status int) (err error) {
	switch status {
	case models.NoRelationship:
		if err = redisutil.RedisHSet(models.RELATION+key, field, "1"); err != nil {
			return
		}
		if err = IncrFollowCount(key); err != nil {
			return
		}
		if err = redisutil.RedisHSet(models.RELATION+field, key, "2"); err != nil {
			return
		}
		if err = IncrFansCount(field); err != nil {
			return
		}
	case models.Follow:
		if err = redisutil.RedisHSet(models.RELATION+key, field, "0"); err != nil {
			return
		}
		if err = DecrFollowCount(key); err != nil {
			return
		}
		if err = redisutil.RedisHSet(models.RELATION+field, key, "0"); err != nil {
			return
		}
		if err = DecrFansCount(field); err != nil {
			return
		}
	case models.Fans:
		if err = redisutil.RedisHSet(models.RELATION+key, field, "3"); err != nil {
			return
		}
		if err = IncrFollowCount(key); err != nil {
			return
		}
		if err = redisutil.RedisHSet(models.RELATION+field, key, "3"); err != nil {
			return
		}
		if err = IncrFansCount(field); err != nil {
			return
		}
	case models.FocusOnEachOther:
		if err = redisutil.RedisHSet(models.RELATION+key, field, "2"); err != nil {
			return
		}
		if err = DecrFollowCount(key); err != nil {
			return
		}
		if err = redisutil.RedisHSet(models.RELATION+field, key, "1"); err != nil {
			return
		}
		if err = DecrFansCount(field); err != nil {
			return
		}
	}
	return nil
}

//GetFollowList 	获取关注的人的列表
func GetFollowList(userid string) ([]*models.RelationInfo, error) {
	vals, err := redisutil.RedisHGetAll(models.RELATION + userid)
	if err != nil {
		return nil, err
	}
	var FollowList []*models.RelationInfo
	for field, val := range vals {
		status, _ := strconv.Atoi(val)
		if status == models.Follow || status == models.FocusOnEachOther {
			userId, _ := strconv.Atoi(field)
			userInfo, _ := GetAUserInfoByUserId(userId)
			RelationInfo := &models.RelationInfo{
				Status:   status,
				UserInfo: userInfo,
			}
			FollowList = append(FollowList, RelationInfo)
		}
	}
	return FollowList, nil
}

//GetFansList 	获取个人粉丝的列表
func GetFansList(userid string) ([]*models.RelationInfo, error) {
	vals, err := redisutil.RedisHGetAll(models.RELATION + userid)
	if err != nil {
		return nil, err
	}
	var FansList []*models.RelationInfo
	for field, val := range vals {
		status, _ := strconv.Atoi(val)
		if status == models.Fans || status == models.FocusOnEachOther {
			userid, _ := strconv.Atoi(field)
			userInfo, _ := GetAUserInfoByUserId(userid)
			RelationInfo := &models.RelationInfo{
				Status:   status,
				UserInfo: userInfo,
			}
			FansList = append(FansList, RelationInfo)
		}
	}
	return FansList, nil
}

//GetFocusOnEachOtherList     获取互相关注的列表、
func GetFocusOnEachOtherList(userid string) ([]*models.RelationInfo, error) {
	vals, err := redisutil.RedisHGetAll(models.RELATION + userid)
	if err != nil {
		return nil, err
	}
	var FansList []*models.RelationInfo
	for field, val := range vals {
		status, _ := strconv.Atoi(val)
		if status == models.FocusOnEachOther {
			userId, _ := strconv.Atoi(field)
			userInfo, _ := GetAUserInfoByUserId(userId)
			RelationInfo := &models.RelationInfo{
				Status:   status,
				UserInfo: userInfo,
			}
			FansList = append(FansList, RelationInfo)
		}
	}
	return FansList, nil
}

//FindFriend   找朋友
func FindFriend() {

}

//获取关注和粉丝总数 Hvals
