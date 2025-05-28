package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"onestep/internal/infrastructure/config"
	"onestep/internal/infrastructure/po"
	"sync"
	"time"
)

// databaseName is the database name of the workspace
var databaseName = "onestep"
var (
	db   *gorm.DB
	once sync.Once
)

// InitDB database, if the ONE-STEP database is not exist, create it.
func InitDB() {
	// create database if not exist
	createDatabaseIfNotExist()
	// initialize tables if not exist
	createTablesIfNotExist()
}

// createTablesIfNotExist creates tables if not exist.
func createTablesIfNotExist() {
	db := GetDB()
	_ = db.AutoMigrate(&po.WorkspacePO{})
}

// GetDB gets a single database connection.
func GetDB() *gorm.DB {
	var databaseConfig = config.AppConfig.Database
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			databaseConfig.User, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseName)
		var err error
		db, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			panic(err)
		}
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.SetMaxIdleConns(databaseConfig.MaxIdleConn)
		sqlDB.SetMaxOpenConns(databaseConfig.MaxOpenConn)
		sqlDB.SetConnMaxLifetime(time.Duration(databaseConfig.MaxLifetime) * time.Second)

	})
	return db
}

// createDatabase if not exist, create it
func createDatabaseIfNotExist() {
	var databaseConfig = config.AppConfig.Database
	// connect database
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=utf8mb4&parseTime=True&loc=Local", databaseConfig.User, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port)))
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// test connection
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}
	var count int
	err = sqlDB.QueryRow("SELECT COUNT(*) FROM information_schema.schemata WHERE schema_name = ?", databaseName).Scan(&count)

	if err != nil {
		panic(err)
	}
	if count == 0 {
		// 数据库不存在，创建它
		_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", databaseName))
		if err != nil {
			panic(err)
		}
	}
	_ = sqlDB.Close()
}
