package mysqlutil

import (
	"CampusDiaryProject/models"
	"CampusDiaryProject/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open("mysql", dsn)

	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func MysqlClose() {
	DB.Close()
}

//InitTable 初始化表模型，包括外键等约束
func InitTable() {
	// 模型绑定
	DB.AutoMigrate(&models.User{}, &models.UserInfo{}, &models.ArticleTypeDict{}, &models.ArticleInfo{}, &models.FollowFansCount{} ,&models.Weight{},&models.Comment{})
	//外键添加
	DB.Model(&models.UserInfo{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	DB.Model(&models.ArticleInfo{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade")
	DB.Model(&models.FollowFansCount{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	DB.Model(&models.ArticleInfo{}).AddForeignKey("article_type_id", "article_type_dicts(id)", "cascade", "cascade")
	DB.Model(&models.Weight{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	DB.Model(&models.Weight{}).AddForeignKey("article_type_id", "article_type_dicts(id)", "cascade", "cascade")
	DB.Model(&models.Comment{}).AddForeignKey("article_id","article_infos(id)","cascade", "cascade")
	DB.Model(&models.Comment{}).AddForeignKey("user_id","users(id)","cascade", "cascade")
}
