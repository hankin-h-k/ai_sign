package service

import (
	Config "github.com/hankin-h-k/ai_sign/config"
)

var initStatus bool

func InitService(host, appId, privateKey string) {
	if !initStatus {
		config := Config.GetInitConfig()
		config.SetHost(host)
		config.SetAppId(appId)
		config.SetPrivateKey(privateKey)
		initStatus = true
	}

}
