package service

import (
	"CampusDiaryProject/models"
	"CampusDiaryProject/utils/mysqlutil"
)

//AddUserInfo 创建个人信息
func AddUserInfo(userInfo *models.UserInfo) (err error) {
	err = mysqlutil.DB.Create(userInfo).Error
	return
}

//UpdateUserInfo 更新用户信息
func UpdateUserInfo(userInfo *models.UserInfo) (err error) {
	err = mysqlutil.DB.Save(userInfo).Error
	return
}

//GetAllUserInfo 获取所有用户信息
func GetAllUserInfo() (userInfoList []*models.UserInfo, err error) {
	if err = mysqlutil.DB.Select("*").
		Joins("inner join follow_fans_counts on user_infos.user_id=follow_fans_counts.user_id").
		Find(&userInfoList).Error; err != nil {
		return nil, err
	}
	return
}

//GetAUserInfoByUserId 根据用户id获取一个用户信息
func GetAUserInfoByUserId(userid int) (userInfo *models.UserInfo, err error) {
	userInfo = new(models.UserInfo)
	if err = mysqlutil.DB.Select("*").Where("user_infos.user_id=?", userid).
		Joins("inner join follow_fans_counts on user_infos.user_id=follow_fans_counts.user_id").
		First(userInfo).Error; err != nil {
		return nil, err
	}
	return
}

//GetAUserNameByUserId 根据用户id获取一个用户名
func GetAUserNameByUserId(userid int) (userName string, err error) {
	var userinfo = models.UserInfo{}
	if err = mysqlutil.DB.Select("user_name").Where("user_id=?", userid).First(&userinfo).Error; err != nil {
		return "", err
	}
	userName = userinfo.UserName
	return
}

//UpdateHeadPhoto  更新头像信息
func UpdateHeadPhoto(filePath string, id int) (err error) {
	err = mysqlutil.DB.Model(&models.UserInfo{}).Where("user_id=?", id).Update("head_photo", filePath).Error
	return
}

//UpdateBackPhoto  更新背景图片信息
func UpdateBackPhoto(filePath string, id int) (err error) {
	err = mysqlutil.DB.Model(&models.UserInfo{}).Where("user_id=?", id).Update("back_photo", filePath).Error
	return
}
