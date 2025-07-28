package model

import (
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, log *log.Logger) {
	if err := db.AutoMigrate(
		&User{},
	); err != nil {
		log.Error("collect migrate error", zap.Error(err))
		panic(err)
	}
	return
}
