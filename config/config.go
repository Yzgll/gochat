package config

import (
	"gochat/utils"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	System System `yaml:"system"`
	Mysql  Mysql  `yaml:"mysql"`
}

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Mysql struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Dbname      string `yaml:"dbname"`
	Maxidle     int    `yaml:"maxidle"`
	Maxopen     int    `yaml:"maxopen"`
	Maxlifetime int    `yaml:"maxlifetime"`
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
