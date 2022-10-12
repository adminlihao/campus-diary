package controller

import (
	"CampusDiaryProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

//UpdateUserInfo 更新用户信息
func UpdateUserInfo(c *gin.Context) {
	userid := c.GetInt("userid")
	userInfo, err := service.GetAUserInfoByUserId(userid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(userInfo)
	if err = service.UpdateUserInfo(userInfo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

//GetAUserInfoByUserId 获取用户信息
func GetAUserInfoByUserId(c *gin.Context) {
	userid := c.GetInt("userid")
	userInfo, err := service.GetAUserInfoByUserId(userid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	user, err := service.GetAUserById(strconv.Itoa(userid))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   map[string]interface{}{"user": map[string]string{"number":user.Number,"email":user.Email}, "userinfo": userInfo},
	})

}

//GetAllUserInfo 获取所有用户信息
func GetAllUserInfo(c *gin.Context) {
	// 查询user_infos这个表里的所有数据
	userInfoList, err := service.GetAllUserInfo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   userInfoList,
		})
	}
}

//HeadPhotoUpload 头像上传
func HeadPhotoUpload(c *gin.Context) {
	userid := c.GetInt("userid")
	//获取表单file的数据
	fileHeader, err := c.FormFile("head_photo_file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".jpeg" || fileExt == ".webp" || fileExt == ".bmp" {
		fileExt = ".jpeg"
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "图片格式错误，请重新上传！"})
		return
	}
	fileDir := "./upload/head_photo/"
	filePath := fileDir + "head_photo_user_" + strconv.Itoa(userid) + fileExt
	err = c.SaveUploadedFile(fileHeader, filePath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		//持久化头像路径
		err = service.UpdateHeadPhoto(filePath, userid)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":     "success",
			"head_photo": filePath,
		})
	}
}

//BackPhotoUpload 背景图片上传
func BackPhotoUpload(c *gin.Context) {
	userid := c.GetInt("userid")
	//获取表单file的数据
	fileHeader, err := c.FormFile("back_photo_file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".jpeg" || fileExt == ".webp" || fileExt == ".bmp" {
		fileExt = ".jpeg"
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "图片格式错误，请重新上传！"})
		return
	}
	fileDir := "./upload/back_photo/"
	filePath := fileDir + "back_photo_user_" + strconv.Itoa(userid) + fileExt
	err = c.SaveUploadedFile(fileHeader, filePath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		//持久化背景图片路径
		err = service.UpdateBackPhoto(filePath, userid)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":     "success",
			"back_photo": filePath,
		})
	}

}
