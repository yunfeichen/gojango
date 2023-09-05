/*
 * @Author: huangchangqing
 * @Date: 2021-07-20 17:01:50
 * @LastEditTime: 2021-07-30 15:13:13
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: mysqldb.go
 */
package utils

import (
	"fmt"
	"gjango/utils/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() error {
	ldb, err := NewDB()
	if err != nil {
		return err
	}
	db = ldb
	logger.Debug("utils.NewDB successfully")
	return nil
}

func CloseDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// NewDB 新建DB连接
func NewDB() (*gorm.DB, error) {
	//TODO:需要之后放到配置文件里
	dsn := GetConfig().Storage.Dbpath
	//dsn := "root:root@tcp(10.10.30.60:3306)/goechoexample?charset=utf8mb4&parseTime=True&loc=Local"
	// 打印数据库连接串
	logger.Debug("打开连接（MySQL）：", dsn)

	// 打开数据库链接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),

		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键
		QueryFields:                              true, // 解决查询全部字段可能不走索引的问题
	})
	// 错误退出
	if err != nil {
		message := fmt.Sprintf("数据库连接异常：%s", err.Error())
		logger.Error(message)
		panic(message)
	}

	// 设置数据库连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)                                 // 空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                                // 最大连接数量
	sqlDB.SetConnMaxIdleTime(time.Minute * time.Duration(60)) // 连接最大可复用时间
	logger.Debug("数据库连接已经建立:", dsn)
	// 获取数据库连接
	return db, nil
}

// GetDB 获取DBhandle
func GetDB() *gorm.DB {
	if db == nil {
		panic("db is nil")
	}
	return db
}
