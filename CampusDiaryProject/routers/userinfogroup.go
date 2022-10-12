package routers

import (
	"CampusDiaryProject/controller"
	"CampusDiaryProject/middleware"
	"github.com/gin-gonic/gin"
)

func UserInfoGroups(r *gin.Engine) {
	UserInfoGroup := r.Group("/userInfo", middleware.JWTAuthMiddleware())
	{
		//更新个人账号
		UserInfoGroup.PUT("/updateUserAccount",controller.UpdateUserAccount)
		//更新个人密码
		UserInfoGroup.PUT("/updateUserPassword",controller.UpdateUserPassword)
		//修改个人信息
		UserInfoGroup.PUT("/updateUserInfo", controller.UpdateUserInfo)
		//查询个人信息
		UserInfoGroup.GET("/getUserInfo", controller.GetAUserInfoByUserId)
		//上传用户头像
		UserInfoGroup.POST("/headPhotoUpload", controller.HeadPhotoUpload)
		//上传背景图片
		UserInfoGroup.POST("/backPhotoUpload", controller.BackPhotoUpload)
	}
}
