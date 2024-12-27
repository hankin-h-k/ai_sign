package service

import (
	"log"

	"github.com/hankin-h-k/ai_sign/config"
	"github.com/hankin-h-k/ai_sign/httpUtils"
)

// 创建待签署文件
func CreateFile(dataJson string, fileConfig *config.FileConfig) []byte {
	apiUrl := "contract/createContract"
	log.Println("创建待签署文件")
	response, err := httpUtils.SendRequest(apiUrl, dataJson, fileConfig)
	if err != nil {
		log.Println("创建待签署文件失败")

	}
	return response
}

// 添加签署方，入参为数组
func AddSigner(dataJson string) []byte {
	apiUrl := "contract/addSigner"
	log.Println("添加签署方")
	response, err := httpUtils.SendRequest2(apiUrl, dataJson, nil)
	if err != nil {
		log.Println("添加签署方失败")

	}
	return response
}

// 下载合同
func DownloadFile(dataJson string) []byte {
	apiUrl := "contract/downloadContract"
	log.Println("下载合同")
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		log.Println("下载合同失败")

	}
	return response
}
