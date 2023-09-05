/*
 * @Author: huangchangqing
 * @Date: 2021-07-20 17:01:50
 * @LastEditTime: 2021-07-30 15:13:13
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: rediscache.go
 */
package utils

import (
	"errors"
	"gjango/utils/logger"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func InitRedisCache() error {
	lrdb, err := NewRedisCache()
	if err != nil {
		return err
	}
	rdb = lrdb

	return nil
}

func NewRedisCache() (*redis.Client, error) {
	redisAddr := GetConfig().Redis.Addr
	redisDB := GetConfig().Redis.DB

	cache := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",      // no password set
		DB:       redisDB, // use default DB
	})
	if cache == nil {
		err := errors.New("初始化Redis连接失败！")
		return nil, err
	}
	logger.Info("已建立Redis缓存")
	return cache, nil
}

func CloseRedisCache() error {
	if rdb != nil {
		return rdb.Close()
	}
	return nil
}
