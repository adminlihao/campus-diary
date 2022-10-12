package models

//User 用户结构体
type User struct {
	ID       int    `json:"userid"`                        //用户id
	Number   string `json:"number" gorm:"unique;not null"` //电话号
	Email    string `json:"email" gorm:"unique;not null"`  //电子邮箱
	Password string `json:"password"`                      //密码
}
