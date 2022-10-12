package routers

import (
	"CampusDiaryProject/controller"
	"CampusDiaryProject/middleware"
	"github.com/gin-gonic/gin"
)

func NotifyGroups(r *gin.Engine) {
	NotifyGroup := r.Group("/notify", middleware.JWTAuthMiddleware())
	{
		//获取通知历史消息
		NotifyGroup.GET("/getNotifyHistory", controller.GetNotifyHistory)
		//获取未读通知消息数量
		NotifyGroup.GET("/getUnreadNotify", controller.GetUnreadNotify)
		//重置未读通知消息数量
		NotifyGroup.PUT("/resetUnreadNotify", controller.ResetUnreadNotify)
		//获取未读聊天信息通知
		NotifyGroup.GET("/getUnreadChat", controller.GetUnreadChat)
		//重置未读聊天信息数目
		NotifyGroup.PUT("/resetUnreadChat", controller.ResetUnreadChat)
	}
}
