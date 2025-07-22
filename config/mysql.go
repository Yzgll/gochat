package config

import (
	"fmt"
	"gochat/migration"
	"gochat/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	//从配置文件获取数据库信息
	var err error
	mysqlconfig := GlobalConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlconfig.User, mysqlconfig.Password, mysqlconfig.Host, mysqlconfig.Port, mysqlconfig.Dbname)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Log.Error("连接数据库失败:%v", err)
		return nil
	}
	sqldb, err := DB.DB()
	if err != nil {
		utils.Log.Error("获取sqldb失败:%v", err)
		return nil
	}
	sqldb.SetMaxIdleConns(mysqlconfig.Maxidle)
	sqldb.SetMaxOpenConns(mysqlconfig.Maxopen)
	sqldb.SetConnMaxLifetime(time.Second * time.Duration(mysqlconfig.Maxlifetime))
	err = sqldb.Ping()
	if err != nil {
		utils.Log.Error("MySQL数据库连接失败！ ")
		return nil
	}
	// 执行自动迁移
	if err := migration.AutoMigrate(DB); err != nil {
		utils.Log.Error("MySQL数据库迁移失败!")
	}
	utils.Log.Infof("MySQL数据库连接成功！ ")
	return DB
}
