package routers

import (
	"CampusDiaryProject/controller"
	"CampusDiaryProject/middleware"
	"github.com/gin-gonic/gin"
)

func ManageGroups(r *gin.Engine) {
	ManageGroup := r.Group("/manage",middleware.JWTAuthMiddleware())
	{
		// 查看所有用户
		ManageGroup.GET("/getAllUser", controller.GetAllUser)
		//查询一个用户
		ManageGroup.GET("/getUser", controller.GetAUserById)
		// 修改某一个用户
		//ManageGroup.PUT("/updateUserAccount", controller.UpdateAUserAccount)
		// 删除某一个用户
		ManageGroup.DELETE("/deleteUser", controller.DeleteAUser)
		//查询所有用户个人信息
		ManageGroup.GET("/getAllUserInfo", controller.GetAllUserInfo)
	}

}
