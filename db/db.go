package db

import (
	"fmt"
	"log"
	"sync"

	"github.com/XiaoBeiAjie/aicongyou/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
    db   *gorm.DB
    once sync.Once
)

func Init() error {
    var initErr error
    once.Do(func() {
        dsn := config.Global.Database.GetDSN()
        gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
            Logger: logger.Default.LogMode(logger.Info),
        })
        if err != nil {
            initErr = fmt.Errorf("failed to connect database: %w", err)
            return
		}
        db = gormDB
        log.Println("Database connection initialized successfully")
    })

    return initErr
}

func GetDB() *gorm.DB {
    if db == nil {
        log.Fatal("Database not initialized. Call database.Init() first")
    }
    return db
}

func Close() error {
    if db != nil {
        sqlDB, err := db.DB()
        if err != nil {
            return err
        }
        return sqlDB.Close()
    }
    return nil
}