package service

import (
	"encoding/json"
	"log"

	Config "github.com/hankin-h-k/ai_sign/config"
	"github.com/hankin-h-k/ai_sign/httpUtils"
)

// 创建待签署文件
func CreateFile(dataJson string, fileConfig *Config.FileConfig) []byte {
	apiUrl := "contract/createContract"
	log.Println("创建待签署文件")
	response, err := httpUtils.SendRequest(apiUrl, dataJson, fileConfig)
	if err != nil {
		log.Println("创建待签署文件失败")

	}
	return response
}

// 添加签署方，入参为数组
func AddSignerJson(dataJson string) []byte {
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

// 查看文件状态
func QueryFileInfo(req *Config.QueryFileInfoRequest) (*Config.QueryFileInfoResponse, error) {
	log.Println("查看文件状态")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "contract/get/upload/file/info"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.QueryFileInfoResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 创建待签署文件
func CreateContract(req *Config.CreateContractRequest) (*Config.CreateContractResponse, error) {
	log.Println("创建待签署文件")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "contract/createContract"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.CreateContractResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 添加签署方
func AddSigner(req []*Config.AddSignerRequest) (*Config.AddSignerResponse, error) {
	log.Println("添加签署方")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "contract/addSigner"
	response, err := httpUtils.SendRequest2(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	// fmt.Println(string(response))
	var resp Config.AddSignerResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 查询合同状态
func QueryContractStatus(req *Config.QueryContractStatusRequest) (*Config.QueryContractStatusResponse, error) {
	log.Println("查询合同状态")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "contract/status"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.QueryContractStatusResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 查询合同信息
func QueryContract(req *Config.QueryContractRequest) (*Config.QueryContractResponse, error) {
	log.Println("查询合同信息")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "contract/getContract"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.QueryContractResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 预览合同
func PreviewContract(req []*Config.PreviewContractRequest) (*Config.PreviewContractResponse, error) {
	log.Println("预览合同")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "contract/previewContract"
	response, err := httpUtils.SendRequest2(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.PreviewContractResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 下载合同
func DownloadContract(req *Config.DownloadContractRequest) (*Config.DownloadContractRespnse, error) {
	log.Println("下载合同")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "contract/downloadContract"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.DownloadContractRespnse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 下载指定合同附件
func DownloadContractByAttach(req *Config.DownloadContractByAttachRequest) (*Config.DownloadContractByAttachRespnse, error) {
	log.Println("下载指定合同附件")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "contract/downloadByAttachName"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.DownloadContractByAttachRespnse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 撤销合同
func WithdrawContract(req *Config.WithdrawContractRequest) (*Config.WithdrawContractResponse, error) {
	log.Println("合同撤销")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "contract/withdraw"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.WithdrawContractResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 作废合同
// 撤销合同
func CancellationContract(req *Config.CancellationContractRequest) (*Config.CancellationContractResponse, error) {
	log.Println("合同作废")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "contract/cancellation"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.CancellationContractResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}

// 延长合同
func DelayContract(req *Config.DelayContractRequest) (*Config.DelayContractResponse, error) {
	log.Println("延长签署过期时间")
	var dataJson string
	if data, err := json.Marshal(req); err == nil {
		dataJson = string(data)
	}
	apiUrl := "contract/delayValidityTime"
	response, err := httpUtils.SendRequest(apiUrl, dataJson, nil)
	if err != nil {
		return nil, err

	}
	var resp Config.DelayContractResponse
	json.Unmarshal(response, &resp)
	return &resp, nil
}
