package test

import (
	"CampusDiaryProject/service"
	"CampusDiaryProject/utils/mysqlutil"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"testing"
)



func Test(t *testing.T)  {
	t.Run("连接数据库", testInit)
	//t.Run("测试分页获取文章资源", testGetAllArticleByPage)
	//t.Run("初始化权重表",testAddWeightInfo)
	//t.Run("权重加1",testWeightIncr)
	//t.Run("测试ffmpeg",testGetCover)
	t.Run("测试查询用户",testSearchUser)
}

func testSearchUser(t *testing.T) {
	userInfoList, err := service.SearchUser("熊",1,8)
	if err != nil {
		log.Println(err)
	}
	for _, info := range userInfoList {
		fmt.Printf("%+v\n",info)
	}
}

func testGetCover(t *testing.T){
	service.GetCover("E:/test.mp4","./cover.jpg")
}

func testInit(t *testing.T){
	mysqlutil.DB,_=gorm.Open("mysql","root:root@tcp(localhost:3306)/redbookproject")

}

func testGetAllArticleByPage(t *testing.T){
	var userid=1
	var pageSize=8
	var pageIndex=1
	dates,_:= service.GetAllArticleByPage(userid,pageIndex,pageSize)
	for i, date := range dates {
		fmt.Println(i,date)
	}
}

func testAddWeightInfo(t *testing.T){
	_ = service.AddWeightInfo(2)
}

func testGetWeightInfo(t *testing.T)  {
	weightList,err:= service.GetWeightInfo(1)
	if err!=nil {
		fmt.Println(err)
		return
	}
	for _, v := range weightList {
		fmt.Println(v)
	}
}

func testWeightIncr(t *testing.T) {
	s := service.WeightDecr("1", 1)
	fmt.Println(s)
}

