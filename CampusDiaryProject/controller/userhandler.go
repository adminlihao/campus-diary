package controller

import (
	"CampusDiaryProject/models"
	"CampusDiaryProject/service"
	"CampusDiaryProject/utils/emails"
	"CampusDiaryProject/utils/jwt"
	"CampusDiaryProject/utils/pwdMgr"
	"CampusDiaryProject/utils/redisutil"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//SendCode 验证码发送
func SendCode(c *gin.Context) {
	email := c.Query("email")
	code, err := emails.SendCode(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	err = redisutil.RedisSet(email, code, 60)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})

}

//CheckCode 判断验证码是否正确
func CheckCode(c *gin.Context) {
	jsonData := make(map[string]string)
	c.BindJSON(&jsonData)
	val, err := redisutil.RedisGet(jsonData["email"])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if val == jsonData["code"] {
		c.JSON(http.StatusOK, gin.H{
			"status":       "success",
			"check_result": true,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"check_result": false,
	})
}

// Register 注册
func Register(c *gin.Context) {
	// 前端页面填写用户信息 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	jsonData := map[string]string{}
	c.BindJSON(&jsonData)
	var user = models.User{
		Number:   jsonData["number"],
		Email:    jsonData["email"],
		Password: pwdMgr.Encryption(jsonData["password"]),
	}
	//2.判定邮箱验证码
	val, err := redisutil.RedisGet(user.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if val == jsonData["code"] {
		// 3. 存入数据库
		err = service.AddUser(&user)
		var userInfo = models.UserInfo{
			UserName: "用户" + user.Number[5:],
			UserID:   user.ID,
		}
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		} else if err = service.AddUserInfo(&userInfo); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		} else if err = service.AddFollowCount(user.ID); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		} else if err = service.AddWeightInfo(user.ID); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
			})
		}
	} else {
		c.JSON(http.StatusOK, "code false")
	}
}

//Login 登录
func Login(c *gin.Context) {
	account := c.PostForm("account")
	pwd := c.PostForm("password")
	user, err := service.CheckAccount(account)
	isPwd := pwdMgr.Decryption(user.Password, pwd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		if isPwd {
			tokenString, _ := jwt.GenToken(user.ID)
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"token":  tokenString,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "failure",
			})
		}
	}
}

//CheckAccount 检查账号
func CheckAccount(c *gin.Context) {
	account := c.Query("account")
	user, err := service.CheckAccount(account)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		if user.ID > 0 {
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
			})
		}
	}
}

//GetAllUser 查询所有用户
func GetAllUser(c *gin.Context) {
	// 查询users这个表里的所有数据
	userList, err := service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   userList,
		})
	}
}

//GetAUserById 根据id获取用户
func GetAUserById(c *gin.Context) {
	userid := c.Query("userid")

	user, err := service.GetAUserById(userid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   user,
		})
	}
}

//UpdateUserAccount 更新账号
func UpdateUserAccount(c *gin.Context) {
	userid := c.GetInt("userid")
	jsonData := make(map[string]string)
	c.BindJSON(&jsonData)

	//2.判定邮箱验证码
	val, err := redisutil.RedisGet(jsonData["email"])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if val == jsonData["code"] {
		if err = service.UpdateUserAccount(userid, jsonData["account_type"], jsonData["account"]); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusOK, "code false")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//UpdateUserPassword 更新密码
func UpdateUserPassword(c *gin.Context) {
	userid := c.GetInt("userid")
	jsonData := make(map[string]string)
	c.BindJSON(&jsonData)
	user, err := service.GetAUserById(strconv.Itoa(userid))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	//2.判定邮箱验证码
	val, err := redisutil.RedisGet(user.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if val == jsonData["code"] {
		if pwdMgr.Decryption(user.Password,jsonData["old_pwd"]){
			err := service.UpdateUserPassword(userid, pwdMgr.Encryption(jsonData["new_pwd"]))
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
		}else {
			c.JSON(http.StatusOK, "old_pwd false")
			return
		}
	} else {
		c.JSON(http.StatusOK, "code false")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//DeleteAUser 删除用户
func DeleteAUser(c *gin.Context) {
	userid := c.Query("userid")
	if err := service.DeleteAUser(userid); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}
