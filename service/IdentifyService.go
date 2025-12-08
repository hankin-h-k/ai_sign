package service

import (
	"encoding/json"
	"log"

	Config "github.com/hankin-h-k/ai_sign/config"
	"github.com/hankin-h-k/ai_sign/httpUtils"
)

func GetAuthPersonIdentifyUrl(req *Config.PersonIdentifyUrlRequest) (*Config.PersonIdentifyUrlResponse, error) {
	log.Println("个人实名认证网页版--------")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "auth/person/identifyUrl"

	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.PersonIdentifyUrlResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

func GetAuthRecordInfo(req *Config.AuthRecordInfoRequest) (*Config.AuthRecordInfoResponse, error) {

	log.Println("查询实名认证详细信息--------")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "auth/getAuthRecordInfo"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.AuthRecordInfoResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}
