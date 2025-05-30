package service

import (
	"encoding/json"
	"log"

	Config "github.com/hankin-h-k/ai_sign/config"

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

// 创建个人账号2.0
func AddPersonalUser(req *Config.CreatePersonalAccountRequest) (*Config.CreatePersonalAccountResponse, error) {
	log.Println("创建个人账号--------")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "v2/user/addPersonalUser"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.CreatePersonalAccountResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
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

// 查询个人账号2.0
func QueryPersonalUser(req *Config.QueryPersonalAccountRequest) (*Config.QueryPersonalAccountResponse, error) {
	log.Println("查询个人账号--------")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "user/getUser"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.QueryPersonalAccountResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 修改用户
func UpdatePersonalUser(req *Config.UpdatePersonalAccountRequest) (*Config.UpdatePersonalAccountResponse, error) {
	log.Println("修改个人账号--------")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "v2/user/modifyUserName"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.UpdatePersonalAccountResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 删除用户
func DeletePersonalUser(req *Config.DeletePersonalAccountRequest) (*Config.DeletePersonalAccountResponse, error) {
	log.Println("修改个人账号--------")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "user/remove"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.DeletePersonalAccountResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 创建企业账号2.0
func AddEnterpriseUser(req *Config.CreateEnterpriseAccountRequest) (*Config.CreateEnterpriseAccountResponse, error) {
	log.Println("创建企业账号--------")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "v2/user/addEnterpriseUser"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.CreateEnterpriseAccountResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 查询企业账号2.0
func QueryEnterpriseUser(req *Config.QueryEnterpriseAccountRequest) (*Config.QueryEnterpriseAccountResponse, error) {
	log.Println("查询企业账号--------")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "user/getCompUser"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.QueryEnterpriseAccountResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 修改企业用户
func UpdateEnterpriseUser(req *Config.UpdateEnterpriseAccountRequest) (*Config.UpdateEnterpriseAccountResponse, error) {
	log.Println("修改企业账号--------")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "v2/user/modifyCompanyInfo"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.UpdateEnterpriseAccountResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}
