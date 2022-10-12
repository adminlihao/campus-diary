package routers

import (
	"CampusDiaryProject/controller"
	"CampusDiaryProject/middleware"
	"github.com/gin-gonic/gin"
)

func RelationInfoGroups(r *gin.Engine) {
	RelationInfoGroup := r.Group("/relationInfo", middleware.JWTAuthMiddleware())
	{
		//获取用户关系信息
		RelationInfoGroup.GET("/getRelation", controller.GetRelation)
		//更改用户关系信息
		RelationInfoGroup.PUT("/changeRelationStatus", controller.ChangeRelationStatus)
		//获取关注的人的列表
		RelationInfoGroup.GET("/getFollowList", controller.GetFollowList)
		//获取粉丝列表
		RelationInfoGroup.GET("/getFansList", controller.GetFansList)
		//获取关注和粉丝的总数
		RelationInfoGroup.GET("/getFollowAndFansCount",controller.GetFollowAndFansCount)
	}
}
