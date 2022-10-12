package routers

import (
	"CampusDiaryProject/controller"
	"CampusDiaryProject/middleware"
	"CampusDiaryProject/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	//使用跨域中间件
	r.Use(middleware.CrossMiddleware())

	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static/", "./view/static/")
	r.StaticFS("/upload", gin.Dir("./upload/", true))
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("./view/*.html")

	r.GET("/",IndexHandler)

	//发送邮箱验证码
	r.GET("/sendCode", controller.SendCode)
	//注册
	r.POST("/register", controller.Register)
	//登录
	r.POST("/login", controller.Login)
	//检查账号
	r.GET("/check", controller.CheckAccount)
	//检查验证码
	r.POST("/checkCode",controller.CheckCode)

	//websocket聊天
	r.GET("/webSocket", controller.WsHandler)

	//用户信息路由组
	UserInfoGroups(r)

	//注册用户关系路由组
	RelationInfoGroups(r)

	//注册管理员操作路由组
	ManageGroups(r)

	//注册文章路由组
	ArticlesGroups(r)

	//注册评论路由组
	CommentGroups(r)

	//注册通知路由组
	NotifyGroups(r)

	//注册搜索路由组
	SearchGroups(r)

	return r
}
