package service

import (
	"CampusDiaryProject/models"
	"CampusDiaryProject/utils/mysqlutil"
	"github.com/jinzhu/gorm"
	"math/rand"
	"sort"
	"time"
)

//WeightSelector 权重选择器
type WeightSelector struct {
	Weight []int
}

//GetWeightSelector 构造权重选择器
func GetWeightSelector(weight []int) WeightSelector {
	for i := 1; i < len(weight); i++ {
		weight[i] += weight[i-1]
	}
	return WeightSelector{weight}
}

//PickIndex 按照权重随机返回值
func (this *WeightSelector) PickIndex() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := r.Intn(this.Weight[len(this.Weight)-1]+ 1)
	return sort.SearchInts(this.Weight, x)
}

//AddWeightInfo 初始化权重信息
func AddWeightInfo(userId int) error{
	articleTypeList, err := GetAllArticleType()
	if err!=nil{
		return err
	}
	for i:=0;i<len(articleTypeList);i++{
		if err= mysqlutil.DB.Create(&models.Weight{
			UserID: userId,
			ArticleTypeID: articleTypeList[i].ID,
			Weight: 100,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

//WeightIncr 权重值加1
func WeightIncr(userId string,typeId int) (err error)  {
	err = mysqlutil.DB.Model(&models.Weight{}).Where("user_id=? and article_type_id=?", userId,typeId).Update("weight", gorm.Expr("weight + ?", 1)).Error
	return
}

//WeightDecr 权重值减1
func WeightDecr(userId string,typeId int) (err error)  {
	err = mysqlutil.DB.Model(&models.Weight{}).Where("user_id=? and article_type_id=?", userId,typeId).Update("weight", gorm.Expr("weight - ?", 1)).Error
	return
}

//GetWeightInfo 获取用户权重信息列表
func GetWeightInfo(userid int) (weightList []*models.Weight,err error){
	if err= mysqlutil.DB.Where("user_id = ?",userid).Find(&weightList).Error; err!=nil{
		return nil, err
	}
	return
}