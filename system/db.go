package system

import (
	"coc-question-bank/model"
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

const (
	DBPATH = "./coc-question-bank.sql"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(DBPATH), &gorm.Config{})
	if err != nil {
		return nil, errors.New("打开数据库失败: " + err.Error())
	}

	// 迁移 schema
	err = db.AutoMigrate(&model.Subject{})
	if err != nil {
		return nil, errors.New("迁移数据库失败: " + err.Error())
	}

	//设置连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	return db, nil
}
