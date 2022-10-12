package routers

import (
	"CampusDiaryProject/controller"
	"CampusDiaryProject/middleware"
	"github.com/gin-gonic/gin"
)

func ArticlesGroups(r *gin.Engine) {
	ArticlesGroup := r.Group("/article", middleware.JWTAuthMiddleware())
	{
		//发布文章
		ArticlesGroup.POST("/publishedArticle", controller.PublishedArticle)
		//获取用户所有文章
		ArticlesGroup.GET("/getAllArticle", controller.GetAllArticle)
		//分页获取用户所有文章
		ArticlesGroup.GET("/getAllArticleByPage",controller.GetAllArticleByPage)
		//根据ID获取文章信息
		ArticlesGroup.GET("/getArticleById",controller.GetArticleById)
		//文章点赞
		ArticlesGroup.PUT("/giveLike",controller.GiveLike)
		//取消文章赞
		ArticlesGroup.PUT("/delLike",controller.DelLike)
		//判断文章是否点赞
		ArticlesGroup.GET("/isGiveLike",controller.IsGiveLike)
		//获取点赞文章总数
		ArticlesGroup.GET("/getGiveLikeArticleCount",controller.GetGiveLikeArticleCount)
		//获取用户的点赞文章
		ArticlesGroup.GET("/getGiveLikeArticle",controller.GetGiveLikeArticle)
		//获取所有文章类型
		ArticlesGroup.GET("/getAllArticleType",controller.GetAllArticleType)

		//检查视频文章是否存在
		ArticlesGroup.GET("/articleExist",controller.ArticleExist)
		//分块上传文件
		ArticlesGroup.POST("/chunkUpload",controller.ChunkUpload)
		//合并切片
		ArticlesGroup.POST("/mergeChunk",controller.MergeChunk)
		//取消上传
		ArticlesGroup.POST("/cancelChunk",controller.CancelChunk)

		//feed流文章推送
		ArticlesGroup.GET("/pushArticleFeed",controller.PushArticleFeed)
		//具体类型文章推送
		ArticlesGroup.GET("/pushArticleByType",controller.PushArticleByType)
		//分页获取朋友的文章
		ArticlesGroup.GET("/getArticleFromFriend",controller.GetArticleFromFriend)
		//分页获取关注的用户的文章
		ArticlesGroup.GET("/getArticleFromFollow",controller.GetArticleFromFollow)


	}
}
