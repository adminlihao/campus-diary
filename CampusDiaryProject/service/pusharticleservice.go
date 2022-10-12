package service

import (
	"CampusDiaryProject/models"
	"CampusDiaryProject/utils/bloomfilter"
	"CampusDiaryProject/utils/redisutil"
	"strconv"
	"time"
)

//ArticleFeed feed推荐获取article
func ArticleFeed(userid int) (ArticleIDs []int, err error) {
	weightList, err := GetWeightInfo(userid)
	var weights, typeIDs []int
	for _, w := range weightList {
		weights = append(weights, w.Weight)
	}
	//权重选择器
	weightSelector := GetWeightSelector(weights)
	for i := 0; i < 8; i++ {
		time.Sleep(time.Nanosecond)
		index := weightSelector.PickIndex()
		typeIDs = append(typeIDs, index+1)
	}
	//根据用户id创建布隆过滤器
	bloomFilter := bloomfilter.NewBloomFilter(models.BLOOM_FILTER + strconv.Itoa(userid))
	for _, typeId := range typeIDs {
		member, err := redisutil.RedisSRandMember(models.ARTICLE_POOL + strconv.Itoa(typeId))
		if err != nil {
			return nil, err
		}
		imember, _ := strconv.Atoi(member)
		exist, err := bloomFilter.Exist(member)
		if err != nil {
			return nil, err
		}
		if exist == 0 && imember > 0 {
			ArticleIDs = append(ArticleIDs, imember)
			err = bloomFilter.Add(member)
			if err != nil {
				return nil, err
			}

		}
	}
	err = redisutil.RedisExpire(bloomFilter.Key, 180)
	if err != nil {
		return nil, err
	}
	return
}

//PushArticleByTypeID 根据具体类型进行推荐
func PushArticleByTypeID(userid int, articleTypeId int) (ArticleIDs []int, err error) {
	//根据用户id创建布隆过滤器
	bloomFilter := bloomfilter.NewBloomFilter(models.BLOOM_FILTER + strconv.Itoa(userid))
	var count int64 = 8
	members, err := redisutil.RedisSRandMemberN(models.ARTICLE_POOL+strconv.Itoa(articleTypeId), count)
	if err != nil {
		return nil, err
	}
	for _, member := range members {
		imember, _ := strconv.Atoi(member)
		exist, err := bloomFilter.Exist(member)
		if err != nil {
			return nil, err
		}
		if exist == 0 && imember > 0 {
			ArticleIDs = append(ArticleIDs, imember)
			err = bloomFilter.Add(member)
			if err != nil {
				return nil, err
			}

		}
	}

	err = redisutil.RedisExpire(bloomFilter.Key, 180)
	if err != nil {
		return nil, err
	}

	return

}

//PushArticleToFriend 推送文章给朋友
