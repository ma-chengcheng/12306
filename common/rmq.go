package common

import "github.com/micro/go-micro/v2/config"

type RocketMQConfig struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
}

func GetRocketMQConfigFromConsul(config config.Config, path ...string) *RedisConfig {
	redisConfig := &RedisConfig{}
	config.Get(path...).Scan(redisConfig)
	return redisConfig
}
