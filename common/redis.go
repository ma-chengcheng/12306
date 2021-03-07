package common

import "github.com/micro/go-micro/v2/config"

type RedisConfig struct {
	Host     string `json:"host"`
	DB       string `json:"db"`
	Port     int64  `json:"port"`
}

func GetRedisConfigFromConsul(config config.Config, path ...string) *RedisConfig {
	redisConfig := &RedisConfig{}
	config.Get(path...).Scan(redisConfig)
	return redisConfig
}
