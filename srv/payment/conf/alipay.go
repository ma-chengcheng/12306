package conf

import "github.com/micro/go-micro/v2/config"

type AlipayConfig struct {
	NotifyUrl  string `json:"notify_url"`
	AppId      string `json:"app_id"`
	PrivateKey string `json:"private_key"`
	IsProd     bool `json:"is_prod"`
}

func GetAlipayFromConsul(config config.Config, path ...string) *AlipayConfig {
	alipayConfig := &AlipayConfig{}
	config.Get(path...).Scan(alipayConfig)
	return alipayConfig
}
