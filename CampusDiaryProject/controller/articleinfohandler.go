package controller

import (
	"CampusDiaryProject/models"
	"CampusDiaryProject/service"
	"CampusDiaryProject/utils/rabbitmqutil"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//PublishedArticle 发布文章
func PublishedArticle(c *gin.Context) {
	userid := c.GetInt("userid")
	isVideo, _ := strconv.Atoi(c.PostForm("is_video"))
	articleTypeId, _ := strconv.Atoi(c.PostForm("article_type_id"))
	articleFileDir := c.PostForm("articleFileDir")
	coverFilePath := c.PostForm("coverFilePath")
	var err error
	if isVideo == 0 {
		coverFilePath, articleFileDir, err = service.UploadPhotoArticleResource(c)
	} else if isVideo == 1 {
		coverFile, _ := c.FormFile("cover_file")
		if coverFile != nil {
			fileExt := filepath.Ext(coverFile.Filename)
			err = c.SaveUploadedFile(coverFile, strings.TrimSuffix(coverFilePath, filepath.Ext(coverFilePath))+fileExt)
		}
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error is file upload failure!": err.Error()})
		return
	}
	articleInfo := &models.ArticleInfo{
		AuthorID:      userid,
		ReleaseTime:   c.PostForm("release_time"),
		Title:         c.PostForm("title"),
		Content:       c.PostForm("content"),
		IsVideo:       isVideo,
		ResourceDir:   articleFileDir,
		CoverUrl:      coverFilePath,
		ArticleTypeID: articleTypeId,
	}
	err = service.AddArticle(articleInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error()})
		return
	}
	err = service.InsertToArticlePool(articleTypeId, articleInfo.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error()})
		return
	}
	err = service.InsertToFriendArticleList(userid, articleInfo.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error()})
		return
	}
	err = service.InsertToFansArticleList(userid, articleInfo.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})

}

//GetAllArticleType 获取所有文章类型
func GetAllArticleType(c *gin.Context) {
	articleTypeList, err := service.GetAllArticleType()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":          "success",
			"articleTypeList": articleTypeList,
		})
	}
}

//GetAllArticle  获取用户所有文章
func GetAllArticle(c *gin.Context) {
	userid := c.Query("userid")
	var id int
	if userid == "" {
		id = c.GetInt("userid")
	} else {
		id, _ = strconv.Atoi(userid)
	}

	articleInfoList, err := service.GetAllArticle(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":          "success",
			"articleInfoList": articleInfoList,
		})
	}
}

//GetAllArticleByPage  分页获取用户所有文章
func GetAllArticleByPage(c *gin.Context) {
	userid := c.Query("userid")
	var id int
	if userid == "" {
		id = c.GetInt("userid")
	} else {
		id, _ = strconv.Atoi(userid)
	}
	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	articleInfoList, err := service.GetAllArticleByPage(id, pageIndex, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":          "success",
			"articleInfoList": articleInfoList,
		})
	}
}

//GetArticleById 根据id获取文章信息
func GetArticleById(c *gin.Context) {
	articleId := c.Query("article_id")
	id, _ := strconv.Atoi(articleId)
	articleInfo, err := service.GetArticleById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"articleInfo": articleInfo,
	})
}

//GetArticleFromFriend   分页获取朋友的文章
func GetArticleFromFriend(c *gin.Context) {
	userid := c.GetInt("userid")
	pageIndex, _ := strconv.ParseInt(c.Query("pageIndex"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.Query("pageSize"), 10, 64)
	articleInfoList, err := service.GetArticleFromFriend(userid, pageIndex, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":          "success",
			"articleInfoList": articleInfoList,
		})
	}
}

//GetArticleFromFollow   分页获取关注的用户的文章
func GetArticleFromFollow(c *gin.Context) {
	userid := c.GetInt("userid")
	pageIndex, _ := strconv.ParseInt(c.Query("pageIndex"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.Query("pageSize"), 10, 64)
	articleInfoList, err := service.GetArticleFromFollow(userid, pageIndex, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":          "success",
			"articleInfoList": articleInfoList,
		})
	}
}

//GiveLike 文章点赞
func GiveLike(c *gin.Context) {
	userid := c.GetInt("userid")
	articleId := c.Query("article_id")
	giveLikeTime := time.Now().UnixNano()
	err := service.GiveLikeByArticleId(strconv.Itoa(userid), articleId, giveLikeTime)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	//通知作者有用户点赞
	userInfo, err := service.GetAUserInfoByUserId(userid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	IArticleId, _ := strconv.Atoi(articleId)
	articleInfo, err := service.GetArticleById(IArticleId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if userid != articleInfo.AuthorID {
		var giveLikeNotify = &models.GiveLikeNotifyInfo{
			NotifyInfo: models.NotifyInfo{
				SendTime: strconv.FormatInt(giveLikeTime, 10),
				FromUser: userInfo,
				ToUserID: articleInfo.AuthorID,
			},
			ArticleInfo: articleInfo,
		}
		err := rabbitmqutil.RMQ.GiveLikeProducer(giveLikeNotify)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//DelLike 取消文章赞
func DelLike(c *gin.Context) {
	userid := c.GetInt("userid")
	articleId := c.Query("article_id")
	err := service.DelLikeByArticleId(strconv.Itoa(userid), articleId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//IsGiveLike   判断文章是否点赞
func IsGiveLike(c *gin.Context) {
	userid := c.GetInt("userid")
	articleId := c.Query("article_id")
	isGiveLike, err := service.IsGiveLikeByArticleId(strconv.Itoa(userid), articleId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"isGiveLike": isGiveLike,
	})
}

//GetGiveLikeArticleCount  获取点赞文章总数
func GetGiveLikeArticleCount(c *gin.Context) {
	userid := c.Query("userid")
	if userid == "" {
		userid = strconv.Itoa(c.GetInt("userid"))
	}
	likeCount, err := service.GetGiveLikeArticleCount(userid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"like_count": likeCount,
	})
}

//GetGiveLikeArticle  获取用户的点赞文章
func GetGiveLikeArticle(c *gin.Context) {
	userid := c.Query("userid")
	strPageIndex := c.Query("pageIndex")
	strPageSize := c.Query("pageSize")
	pageIndex, _ := strconv.ParseInt(strPageIndex, 10, 64)
	pageSize, _ := strconv.ParseInt(strPageSize, 10, 64)
	if userid == "" {
		userid = strconv.Itoa(c.GetInt("userid"))
	}
	giveLikeArticleList, err := service.GetGiveLikeArticleList(userid, pageIndex, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":              "success",
		"giveLikeArticleList": giveLikeArticleList,
	})
}
