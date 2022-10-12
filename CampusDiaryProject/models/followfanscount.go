package models

type FollowFansCount struct {
	UserID      int `json:"userid" gorm:"primary_key;unique;not null;auto_Increment:false;"` //外键绑定用户id
	FollowCount int `json:"follow_count" gorm:"default:0"`                                   //关注总数
	FansCount   int `json:"fans_count" gorm:"default:0"`                                     //粉丝总数
}
