package config

import (
	"gochat/utils"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	System SystemConfig `yaml:"system"`
	Mysql  MySQLConfig  `yaml:"mysql"`
	Redis  RedisConfig  `yaml:"redis"`
	// Sms    SmsConfig    `yaml:sms`
}

// 阿里云短信配置
// type SmsConfig struct {
// 	Enable          bool   `yaml:"enable"`
// 	AccessKeyID     string `yaml:"access_key_id"`
// 	AccessKeySecret string `yaml:"access_key_secret"`
// 	SignName        string `yaml:"sign_name"`
// 	TemplateCode    string `yaml:"template_code"`
// 	RegionID        string `yaml:"region_id"`
// }

type SystemConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type MySQLConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Dbname      string `yaml:"dbname"`
	Maxidle     int    `yaml:"maxidle"`
	Maxopen     int    `yaml:"maxopen"`
	Maxlifetime int    `yaml:"maxlifetime"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`

	// 连接池核心配置（仅保留3个必选参数）
	PoolSize     int `yaml:"pool_size"`
	MinIdleConns int `yaml:"min_idle_conns"`
	MaxIdleConns int `yaml:"max_idle_conns"`
}

var GlobalConfig Config

func InitConfig() {
	configFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		utils.Log.Error("读取配置文件失败: ")
		return
	}
	//反序列化
	yaml.Unmarshal(configFile, &GlobalConfig)
	utils.Log.Infof("读取配置文件成功")
}
