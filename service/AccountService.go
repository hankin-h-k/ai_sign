package service

import (
	"log"

	"github.com/hankin-h-k/ai_sign/httpUtils"
)

// 创建个人账号
func CreatePersonalAccount(dataJson string) []byte {
	apiUrl := "v2/user/addPersonalUser"
	log.Println("创建个人账号--------")
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		log.Println("请求失败：", err)

	}
	return response
}

// 查询个人账号
func QueryPersonalAccount(dataJson string) []byte {
	apiUrl := "user/getUser"
	log.Println("查询个人账号--------")
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		log.Println("请求失败：", err)
	}
	return response
}
