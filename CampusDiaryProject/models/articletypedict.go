package models


//ArticleTypeDict 文章类型字典
/*
	1    谷类
	2	 豆类
	3    薯类
	4	 水果
	5	 蔬菜
	6	 其他
 */
type ArticleTypeDict struct {
	ID	int  `json:"article_type_id"`
	ArticleType    string   `json:"article_type" gorm:"not null;"`
}


