package service

import (
	"CampusDiaryProject/models"
	"CampusDiaryProject/utils/mysqlutil"
)

//AddUser 添加一个用户
func AddUser(user *models.User) (err error) {
	err = mysqlutil.DB.Create(user).Error
	return
}

//GetAllUser 获取所有用户
func GetAllUser() (userList []*models.User, err error) {
	if err = mysqlutil.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

//GetAUserById 获取一个用户
func GetAUserById(id string) (user *models.User, err error) {
	user = new(models.User)
	if err = mysqlutil.DB.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

//CheckAccount 账号检查
func CheckAccount(account string) (user *models.User, err error) {
	user = new(models.User)
	if err = mysqlutil.DB.Where("number=? or email=?", account, account).First(user).Error; err != nil {
		return nil, err
	}
	return
}

//CheckAccountAndPwd 账号和密码检查
func CheckAccountAndPwd(account string, pwd string) (user *models.User, err error) {
	user = new(models.User)
	if err = mysqlutil.DB.Where("(number=? or email=?) and password=?", account, account, pwd).First(user).Error; err != nil {
		return nil, err
	}
	return
}

//UpdateAUser 更新用户
func UpdateAUser(user *models.User) (err error) {
	err = mysqlutil.DB.Save(user).Error
	return
}

//UpdateUserAccount 更新用户账号
func UpdateUserAccount(userid int,accountType string,account string) (err error) {
	err = mysqlutil.DB.Model(&models.User{}).Where("id = ?",userid).Update(accountType,account).Error
	return
}

//UpdateUserPassword 更新用户密码
func UpdateUserPassword(userid int,newPassword string) (err error) {
	err = mysqlutil.DB.Model(&models.User{}).Where("id = ?",userid).Update("password",newPassword).Error
	return
}

//DeleteAUser 删除一个用户
func DeleteAUser(id string) (err error) {
	err = mysqlutil.DB.Where("id=?", id).Delete(&models.User{}).Error
	return
}
