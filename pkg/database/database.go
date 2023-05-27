package database

import (
	"kratos-blog/internal/conf"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(cfg *conf.Bootstrap) (*gorm.DB, error) {
	c := cfg.Data
	dns := c.Database.Source
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(int(c.Database.MaxIdleConns))
	sqlDB.SetMaxOpenConns(int(c.Database.MaxOpenConns))
	sqlDB.SetConnMaxLifetime(time.Second * 25)

	return db, err
}
