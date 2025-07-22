package migration

import (
	"gochat/model"

	"gorm.io/gorm"
)

// AutoMigrate 自动迁移所有模型
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{}, // 注册用户模型
		// 后续可添加其他模型...
	)
}
