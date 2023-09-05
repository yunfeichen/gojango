/*
 * @Author: huangchangqing
 * @Date: 2021-08-17 17:01:50
 * @LastEditTime: 2021-07-30 15:13:13
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \gechosample\datastore\mysqldb.go
 */

package utils

import (
	"gjango/utils/logger"
	"path/filepath"
	"reflect"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

// 配置文件对应的结构体
type TomlConfig struct {
	Listen struct {
		Port int `toml:"port"`
	}

	Storage struct {
		Dbpath string `toml:"dbpath"`
	}

	Redis struct {
		Addr string `toml:"addr"`
		DB   int    `toml:"db"`
	} `toml:"redis"`

	Nats struct {
		Addr string `toml:"addr"`
	} `toml:"nats"`

	Shin3rd struct {
		Addr string `toml:"addr"`
	} `toml:"shin3rd"`

	StaticFile struct {
		Alias string `toml:"alias"`
		Path  string `toml:"path"`
	} `toml:"staticfile"`

	Jwt struct {
		Key     string `toml:"key"`
		Timeout int    `toml:"timeout"`
	} `toml:"jwt"`

	KeyAuth struct {
		KeyName  string `toml:"keyname"`
		KeyValue string `toml:"keyvalue"`
	} `toml:"keyauth"`

	WechatLogin struct {
		AppId     string `toml:"appid"`
		AppSecret string `toml:"appsecret"`
	} `toml:"wechatlogin"`

	WechatPay struct {
		AppId      string `toml:"appid"`
		MerchantId string `toml:"mchid"`
	} `toml:"wechatpay"`
}

var config TomlConfig
var once sync.Once

func printSectionDetail(stu interface{}) {
	var typeInfo = reflect.TypeOf(stu)
	var valInfo = reflect.ValueOf(stu)
	num := typeInfo.NumField()
	for i := 0; i < num; i++ {
		logger.Debug(typeInfo.Field(i).Name, "=", valInfo.Field(i))
	}
}

func GetConfig() *TomlConfig {
	once.Do(func() { //配置文件只读取一次
		cfgFilePath, err := filepath.Abs("./conf/app.toml")
		if err != nil {
			logger.Debug("未获取到conf/app.toml的绝对路径", err)
			return
		}
		//utils.Debug("parse toml file. filePath: %s\n", cfgFilePath)
		if _, err := toml.DecodeFile(cfgFilePath, &config); err != nil {
			logger.Debug("toml文件解码失败", err)
			return
		}
		if gin.DebugMode == "debug" {
			printSectionDetail(config.Storage)
		}

	})
	return &config
}
