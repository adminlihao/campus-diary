package routers

import (
	"CampusDiaryProject/controller"
	"CampusDiaryProject/middleware"
	"github.com/gin-gonic/gin"
)

func CommentGroups(r *gin.Engine) {
	CommentGroup := r.Group("/comment", middleware.JWTAuthMiddleware())
	{
		//添加评论
		CommentGroup.POST("/addComment", controller.AddComment)
		//是否有删除权限
		CommentGroup.GET("/removePermission", controller.RemovePermission)
		//删除评论
		CommentGroup.DELETE("/deleteComment", controller.DeleteComment)
		//评论点赞
		CommentGroup.PUT("/incrCommentGiveLike", controller.IncrCommentGiveLike)
		//取消评论点赞
		CommentGroup.PUT("/decrCommentGiveLike", controller.DecrCommentGiveLike)
		//判断评论是否点赞
		CommentGroup.GET("/isCommentGiveLike", controller.IsCommentGiveLike)
		//根据热度获一级评论
		CommentGroup.GET("/getCommentByHeat", controller.GetCommentByHeat)
		//根据热度获二级评论
		CommentGroup.GET("/getReplyCommentByHeat", controller.GetReplyCommentByHeat)
		//根据commentId获取评论信息
		CommentGroup.GET("/getCommentById", controller.GetCommentById)
	}
}
