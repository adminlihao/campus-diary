package setting

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var Conf = new(AppConfig)

// AppConfig 应用程序配置
type AppConfig struct {
	Release         bool `yaml:"release"`
	Port            int  `yaml:"port"`
	*MySQLConfig    `yaml:"mysql"`
	*RedisConfig    `yaml:"redis"`
	*RabbitMQConfig `yaml:"rabbitmq"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`   //数据库名
	Host     string `yaml:"host"` //地址
	Port     int    `yaml:"port"` //端口
}

// RedisConfig redis配置
type RedisConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`     //地址
	Port     int    `yaml:"port"`     //端口号
	DB       int    `yaml:"db"`       //redis数据库下标
	PoolSize int    `yaml:"poolSize"` //连接池大小
}

//RabbitMQConfig  rabbitmq配置
type RabbitMQConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"` //地址
	Port     int    `yaml:"port"` //端口号
}

func InitConfig(file string) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return  yaml.Unmarshal(yamlFile, Conf)
}
