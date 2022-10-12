package main

import (
	"CampusDiaryProject/routers"
	"CampusDiaryProject/setting"
	"CampusDiaryProject/utils/mysqlutil"
	"CampusDiaryProject/utils/rabbitmqutil"
	"CampusDiaryProject/utils/redisutil"
	ws "CampusDiaryProject/websocket"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请输入: ./main conf/config.yalm")
		return
	}
	var err error
	// 加载配置文件
	if err = setting.InitConfig(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 创建数据库
	// sql: CREATE DATABASE redbookproject;
	// 连接Mysql数据库
	err = mysqlutil.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysqlutil.MysqlClose() // 程序退出关闭mysql连接

	//初始化MySQL表模型
	mysqlutil.InitTable()

	//连接Redis
	err = redisutil.InitRedis(setting.Conf.RedisConfig)
	if err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redisutil.RedisClose() // 程序退出关闭redis连接

	//连接并初始化RabbitMQ
	err = rabbitmqutil.RMQ.InitRabbitMQ(setting.Conf.RabbitMQConfig)
	if err != nil {
		fmt.Printf("init rabbitmq failed, err:%v\n", err)
		return
	}
	//启动协程监听confirm发布确认
	go rabbitmqutil.RMQ.ListenConfirm()
	//启动消费者
	rabbitmqutil.RMQ.StartConsumers()

	defer rabbitmqutil.RMQ.RabbitMQClose() // 程序退出关闭rabbitmq连接

	//启动ws服务
	go ws.WsManager.Start()

	// 路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
