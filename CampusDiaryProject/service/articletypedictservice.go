package service

import (
	"CampusDiaryProject/models"
	"CampusDiaryProject/utils/mysqlutil"
)

//GetArticleTypeById 根据类型id获取文章类型
func GetArticleTypeById(id int) (articleType *models.ArticleTypeDict) {
	articleType = new(models.ArticleTypeDict)
	if err := mysqlutil.DB.Where("id = ?", id).First(articleType).Error; err != nil {
		return nil
	}
	return
}

//GetAllArticleType  获取所有文章类型
func GetAllArticleType() (articleTypeList []*models.ArticleTypeDict,err error) {
	if err = mysqlutil.DB.Find(&articleTypeList).Error; err != nil {
		return nil,err
	}
	return
}
