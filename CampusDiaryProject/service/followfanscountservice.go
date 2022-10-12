package service

import (
	"CampusDiaryProject/models"
	"CampusDiaryProject/utils/mysqlutil"
	"github.com/jinzhu/gorm"
)


//AddFollowCount 初始化个人关注粉丝列
func AddFollowCount(userid int) (err error) {
	err= mysqlutil.DB.Create(&models.FollowFansCount{
		UserID:      userid,
	}).Error
	return
}

//IncrFollowCount 关注总数+1
func IncrFollowCount(userid string) (err error) {
	err= mysqlutil.DB.Model(&models.FollowFansCount{}).Where("user_id=?",userid).Update("follow_count",gorm.Expr("follow_count + ?",1)).Error
	return
}

//IncrFansCount 粉丝总数+1
func IncrFansCount(userid string) (err error) {
	err= mysqlutil.DB.Model(&models.FollowFansCount{}).Where("user_id=?",userid).Update("fans_count",gorm.Expr("fans_count + ?",1)).Error
	return
}


//DecrFollowCount 关注总数-1
func DecrFollowCount(userid string) (err error) {
	err= mysqlutil.DB.Model(&models.FollowFansCount{}).Where("user_id=?",userid).Update("follow_count",gorm.Expr("follow_count - ?",1)).Error
	return
}

//DecrFansCount 粉丝总数-1
func DecrFansCount(userid string) (err error) {
	err= mysqlutil.DB.Model(&models.FollowFansCount{}).Where("user_id=?",userid).Update("fans_count",gorm.Expr("fans_count - ?",1)).Error
	return
}


//GetFollowFansCount 获取关注和粉丝的总数
func GetFollowFansCount(userid string) (FollowFansCount *models.FollowFansCount,err error) {
	FollowFansCount = new(models.FollowFansCount)
	if err = mysqlutil.DB.Where("user_id=?", userid).First(FollowFansCount).Error; err != nil{
		return nil, err
	}
	return
}
