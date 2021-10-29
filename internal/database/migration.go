package database

import (
	"github.com/gretarst/go-rest-api/internal/comment"
	"github.com/jinzhu/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
