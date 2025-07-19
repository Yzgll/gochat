package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement" json:"id"`        //自增主键，数据库内部使用
	Ggnumber     uint64 `gorm:"uniqueIndex" json:"ggnumber"`               //gg号，类似于qq号码，用于登录,唯一
	Password     string `gorm:"size:255;not null" json:"password"`         //登录密码
	Nickname     string `gorm:"size:50;not null" json:"nickname"`          //用户昵称
	Avatarurl    string `gorm:"size:255;default:''" json:"avatarurl"`      //头像url
	Gender       int8   `gorm:"default:0" json:"gender"`                   //性别 1男 2女
	Birthdate    string `gorm:"size:20;not null" json:"birthdate"`         //生日
	Country      string `gorm:"size:50" json:"country"`                    //国家
	City         string `gorm:"size:50" json:"city"`                       //城市
	Signature    string `gorm:"size:255;default:''" json:"signature"`      //个性签名
	Phone        string `gorm:"uniqueIndex;size:20;not null" json:"phone"` //手机
	Email        string `gorm:"uniqueIndex;size:50;not null" json:"email"` //邮箱
	Registertime int64  `gorm:"autoCreateTime" json:"registertime"`        //注册时间
	Status       int8   `gorm:"default:1" json:"status"`                   //账号状态 0封禁 1正常 2冻结
}

func InitModel(db *gorm.DB) {
	db.AutoMigrate(&User{})

}
