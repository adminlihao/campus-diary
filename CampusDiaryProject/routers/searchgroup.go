package routers

import (
	"CampusDiaryProject/controller"
	"CampusDiaryProject/middleware"
	"github.com/gin-gonic/gin"
)

func SearchGroups(r *gin.Engine) {
	SearchGroup := r.Group("/search", middleware.JWTAuthMiddleware())
	{
		//根据用户名搜索用户
		//SearchGroup.GET("/searchUserByUserName", controller.SearchUserByUserName)
		//根据账号和用户名搜索用户
		SearchGroup.GET("/searchUser", controller.SearchUser)
		//根据标题或内容搜索文章
		SearchGroup.GET("/searchArticle", controller.SearchArticle)
		//根据文章类型以及标题或内容搜索文章
		SearchGroup.GET("/searchArticleByType", controller.SearchArticleByType)
		//删除搜索记录
		SearchGroup.DELETE("/delSearchRecord", controller.DelSearchRecord)
		//删除全部搜索记录
		SearchGroup.DELETE("/delAllSearchRecord", controller.DelAllSearchRecord)
		//获取所有搜索记录
		SearchGroup.GET("/getAllSearchRecord", controller.GetAllSearchRecord)
	}
}
