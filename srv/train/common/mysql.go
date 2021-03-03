package common

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfig struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pwd string `json:"pwd"`
	Database string `json:"database"`
	Port int64 `json:"port"`
}

//获取 mysql 的配置
func GetMysqlConfigFromEtcd(config config.Config,path ...string) *MysqlConfig{
	mysqlConfig := &MysqlConfig{}
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}


func GetMySqlDB() (db *gorm.DB, err error) {
	cfg, err := ini.Load("config.ini")

	if err != nil {
		fmt.Printf("Open MySQL database: %v", err)
	}

	mysqlCfg := cfg.Section("mysql")
	username := mysqlCfg.Key("username").String()
	password := mysqlCfg.Key("password").String()
	address := mysqlCfg.Key("host").String() + ":" + mysqlCfg.Key("port").String()
	dbname := mysqlCfg.Key("dbname").String()

	dsn := username + ":" + password + "@tcp(" + address + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 缓存预编译语句
		PrepareStmt: true,
	})
}
