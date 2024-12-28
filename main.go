package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	Config "github.com/hankin-h-k/ai_sign/config"
	Service "github.com/hankin-h-k/ai_sign/service"
	Utils "github.com/hankin-h-k/ai_sign/utils"
)

func main() {
	var (
		host       = "https://prev.asign.cn/" //测试环境 "https://oapi.asign.cn/" //正式环境
		appId      = "392555288"
		privateKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCuDyBfizF8HF4omrNTt+F34pxrUYmmBP3W2B1tSqZtQKidgdGYSq6Slq3Hj/hvOUkpZUCJy0z2X4NFfoFxQgd2cpsQ2rKJGn/tjKoHYBlW9IBIlcOrVXDjCj01oGHDC7728kfmDkQYw3MWIyr0NH4c1KnJ1787atMVn5ICZwWwtm575XiW7FoJUPA85sQ1EothkG4zr/L40Y2AoNoaIBMYznFdah0QjtSVL1gJEYKl1CN2TdJIVz5+SStfWebYssEtAsrcfASnxANpZIeO6nHIEf4D8l6PLVfyvQYzwLFBpdj6XlJ62l2d57DbJ6/8w9F0H0hBDXaxN8rwlKnisx+JAgMBAAECggEAALWV9drA4Oi2UMdg6PcPNA0XEPwBc2l9uKl9nc1BOcdYdGpiY6XJYEN6yQCh01Ow6x3vOAZ2QUYu44ttIxjVwFuiieqjIVM8zWjWfNemtb9gg+/zRL0ku0Gcf/Mc9m7tj73rUN4P3hdy2YrM/62MRUeCzjb9pIz5c1CzasXUe9jdyM96woBOtBftuugt0YF8dP/3K44u5whPiZmf0GmcGRhn/yy8W4ZYBgOD9fQffEBWD/mTerYNMk0jerf6j0sKy7wbPhWW/VQjibvxaQ7uRYCkyrhoZuGMX4ucx8mTlz0+ed3oUWymnV8c390tko/7/CcGGBMlknDFHidMIbfdEQKBgQDoxHOThWAJfgiqx2a+6gOp0q3xyD/98VbqomwrFThJjvm1AnyR9fVeUZhfzz4iOdZWY56/qkQcRfdb3CemQb6MfHUwgzzSU7xf6iQhVZ5QicqmN/ragFXhQDX2YDUg6LY2K4k6YpUv4/0iaPHSJhWHpYCV/6MNAnkODnqQgKR2fQKBgQC/bpd6A9nm3OW/LFAAjd0bvY3vucUBY9Ofyzwq/W8yaHAblLJHS47ao3WNRbWYpvPhJw2adEcgCSL7R8obzFbHWKywNkv6LBwb04g40PyuaEBi80C+kj6HXTtPKv9uwyMH4WMQbYnoOozYfRgx8QWF2ctehRlPs2MtY6OJDiL+/QKBgQCBL48JwAt2GToMNpEiwlR4K1KZD0/cy9Y2cSDhpGxR/wU93fxvMfgVeem8uxO8ASehWLbhn0kf/AD2PbDPyEtOQ0DzRNM71HeHMDcPPSqZfnA3YnhaTsHbOzz32ZuDAUUnuW/3NIGiS8HRGDemL7bhSwybs362GLpA/+4sRwmALQKBgGxh/+88GPXloUowm4cEyuf21VKh2hFkZJ+3IHSykm92w+pJsPo+pD3TiC2ZByzLcY7LXp9RMCfKJH0icyJz+bnZ0cPeqfuVV8DFbf3FDRrtuW6MjN6YzEA1n+u9y2r5nm7KxBy79/V8mNvFT5qO/HKBEeyM7Py78u3y2NczIss9AoGBAMq2qAKrhI/p71Koy6ci9RRQEZMNwVk+e2RlDNzX0nq+yKzwivTSPWb5wXzjTIqLzgb2P+aAnKZ+Zfs0YIZMdqo5p696Lt2jvRmiiRoe0go101ZnPTxnNSZ/YCjbP49NpPHjEjfeO6IpAOWB7QwvFKS/krHb+fKFkTuyrD5su8PC"
	)
	Service.InitService(host, appId, privateKey)

	// createFile()

	// personalAccount := Config.CreatePersonalAccountRequest{}
	// personalAccount.Account = "local_6"
	// personalAccount.Name = "Hankin"
	// personalAccount.IdCard = "511112200107212811"
	// res, err := Service.AddPersonalUser(&personalAccount)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// if res.Code != 100000 {
	// 	log.Println(res.Msg)
	// } else {
	// 	log.Println(res.Data)
	// }

	// var req Config.QueryPersonalAccountRequest
	// // req.Account = "local_1"
	// res, err := Service.QueryPersonalUser(&req)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// var req Config.UpdatePersonalAccountRequest
	// req.Account = "local_1"
	// req.Name = "Hankin1"
	// req.IdentifyType = 2
	// req.Mobile = "15872844805"

	// res, err := Service.UpdatePersonalUser(&req)

	// var req Config.DeletePersonalAccountRequest
	// req.Account = "local_5"
	// res, err := Service.DeletePersonalUser(&req)

	// var req Config.QueryFileInfoRequest
	// req.FileId = "testContractNo1735282333"
	// res, err := Service.QueryFileInfo(&req)

	// var contractNo = "testContractNo" + strconv.Itoa(int(time.Now().Unix()))
	// var req Config.CreateContractRequest
	// req.ContractNo = contractNo
	// req.ContractName = "合同" + time.Now().Format(time.DateTime)
	// req.ContractNo = contractNo
	// req.ValidityTime = 30
	// req.SignOrder = 1
	// var template1 Config.Templates
	// template1.TemplateNo = "TN79662B1B21B64FB7AA91D7051860EB23"
	// template1.FillData = make(map[string]string)
	// template1.FillData["idcard"] = "511112200107212811"
	// template1.FillData["name"] = "Hankin"
	// template1.FillData["address"] = "广东深圳"

	// pic1 := readFile("https://image.fulllinkai.com/202309/15/99dc4bf7859476aabeeaa83f6fa4a253.png")
	// var componentData1 = map[string]interface{}{}
	// componentData1["type"] = 11
	// componentData1["keyword"] = "card_face"
	// componentData1["imageBase64"] = pic1

	// var componentData2 = map[string]interface{}{}
	// componentData2["type"] = 11
	// componentData2["keyword"] = "card_emblem"
	// pic2 := readFile("https://images.health.ufutx.com/202412/17/9e81af2c8c244e4c3e71b6b2650db24e.jpeg")
	// componentData2["imageBase64"] = pic2
	// template1.ComponentData = append(template1.ComponentData, componentData1, componentData2)

	// req.Templates = append(req.Templates, template1)
	// req.ReadSeconds = 30
	// req.ReadType = 1
	// req.NeedAgree = 1
	// req.NotifyUrl = "https://health.ufutx.com"

	// _, err := Service.CreateContract(&req)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// // if res.Code != 100000 {
	// // 	log.Println(res.Msg)
	// // } else {
	// // 	log.Println(res.Data)
	// // }

	// var signReq Config.AddSignerRequest
	// signReq.ContractNo = contractNo
	// signReq.Account = "local_1"
	// signReq.SignType = 3
	// signReq.ValidateType = 7
	// var signStrategy Config.SignStrategy
	// signStrategy.AttachNo = 1
	// signStrategy.LocationMode = 4
	// signStrategy.SignKey = "甲方"
	// signStrategy.SignPage = 7

	// signReq.SignStrategyList = append(signReq.SignStrategyList, signStrategy)

	// var signReq2 Config.AddSignerRequest
	// signReq2.ContractNo = contractNo
	// signReq2.Account = "测试 Account"
	// signReq2.SignType = 3
	// signReq2.ValidateType = 1
	// var signStrategy2 Config.SignStrategy
	// signStrategy2.AttachNo = 1
	// signStrategy2.LocationMode = 4
	// signStrategy2.SignKey = "乙方"
	// signReq2.SignStrategyList = append(signReq2.SignStrategyList, signStrategy2)

	// var signReqList []*Config.AddSignerRequest
	// signReqList = append(signReqList, &signReq)
	// _, err = Service.AddSignerV2(signReqList)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// var req1 Config.PreviewContractRequest
	// req1.ContractNo = contractNo
	// req1.Account = "local_1"
	// var signStrategy Config.SignStrategy
	// signStrategy.AttachNo = 1
	// signStrategy.LocationMode = 4
	// signStrategy.SignKey = "甲方"
	// signStrategy.SignPage = 7
	// req1.SignStrategyList = append(req1.SignStrategyList, signStrategy)

	// var list []*Config.PreviewContractRequest
	// list = append(list, &req1)
	// Service.PreviewContract(list)

	// 创建用户-start
	// var createPersonal Config.CreatePersonalAccountRequest
	// createPersonal.Account = "local_1"
	// //createPersonal.SerialNo = "***" // 实名认证流水号（若您希望不传此值，由您自行完成认证。可登录开放平台【应用配置】内设置）
	// createPersonal.Name = "黄济"
	// createPersonal.IdCard = "511112200107212811"
	// createPersonal.IdCardType = 1 // 默认传1
	// createPersonal.Mobile = "15872844805"
	// var createPersonalJson string
	// if data, err := json.Marshal(createPersonal); err == nil {
	// 	createPersonalJson = string(data)
	// }
	// var initResult []byte = Service.CreatePersonalAccount(createPersonalJson)
	// CreatePsnByThirdData := Utils.BytetoJson(initResult)["data"]
	// // account := CreatePsnByThirdData.(map[string]interface{})["account"].(string)
	// fmt.Println(CreatePsnByThirdData)
	// 创建用户-end

	// 查询用户信息-start
	// var queryUserInfo Config.QueryPersonalAccountConfig
	// queryUserInfo.Account = "local_1"
	// var queryUserInfoJson string
	// if data, err := json.Marshal(queryUserInfo); err == nil {
	// 	queryUserInfoJson = string(data)
	// }
	// var initResult []byte = Service.QueryPersonalAccount(queryUserInfoJson)
	// CreatePsnByThirdData := Utils.BytetoJson(initResult)["data"]
	// fmt.Println(CreatePsnByThirdData)
	// 查询用户信息 - end

	//添加签署方-end

	// var req Config.DownloadContractRequest
	// req.ContractNo = "testContractNo1735357438"
	// res, _ := Service.DownloadContract(&req)
	// Utils.SaveFileByBase64(res.Data.Data, "D:\\www\\test.pdf")

	// var req Config.WithdrawContractRequest
	// req.ContractNo = "testContractNo1735357438"
	// Service.WithdrawContract(&req)

	////下载合同-start
	// var downloadFileConfig Config.DownloadFileConfig
	// downloadFileConfig.ContractNo = contractNo
	// var downloadFileConfigJson string
	// if data, err := json.Marshal(downloadFileConfig); err == nil {
	// 	downloadFileConfigJson = string(data)
	// }
	// var initResult []byte = Service.DownloadFile(downloadFileConfigJson)
	// CreatePsnByThirdData := Utils.BytetoJson(initResult)["data"]
	// base64String := CreatePsnByThirdData.(map[string]interface{})["data"].(string)
	// Utils.SaveFileByBase64(base64String, "D:\\临时文件\\test.pdf")
	//// 下载合同-end

	// var req Config.DelayContractRequest
	// req.ContractNo = "testContractNo1735357438"
	// req.ValidityDate = "2025-12-30"
	// Service.DelayContract(&req)

}

func createFile() {
	//创建待签署文件(模板)-start

	var contractNo = "testContractNo" + strconv.Itoa(int(time.Now().Unix()))
	var createFile Config.CreateFileConfig
	createFile.ContractName = "第一份合同" + time.Now().Format(time.DateTime)
	createFile.ContractNo = contractNo
	createFile.ValidityTime = 30
	createFile.SignOrder = 1
	var template1 Config.Templates
	template1.TemplateNo = "TN79662B1B21B64FB7AA91D7051860EB23"
	template1.FillData = make(map[string]string)
	template1.FillData["idcard"] = "511112200107212811"
	template1.FillData["name"] = "Hankin"
	template1.FillData["address"] = "广东深圳"

	pic1 := readFile("https://image.fulllinkai.com/202309/15/99dc4bf7859476aabeeaa83f6fa4a253.png")
	var componentData1 = map[string]interface{}{}
	componentData1["type"] = 11
	componentData1["keyword"] = "card_face"
	componentData1["imageBase64"] = pic1

	var componentData2 = map[string]interface{}{}
	componentData2["type"] = 11
	componentData2["keyword"] = "card_emblem"
	pic2 := readFile("https://images.health.ufutx.com/202412/17/9e81af2c8c244e4c3e71b6b2650db24e.jpeg")
	componentData2["imageBase64"] = pic2
	template1.ComponentData = append(template1.ComponentData, componentData1, componentData2)

	createFile.Templates = append(createFile.Templates, template1)

	var createFileJson string
	if data, err := json.Marshal(createFile); err == nil {
		createFileJson = string(data)
	}
	fileResult := Service.CreateFile(createFileJson, nil)
	fileData := Utils.BytetoJson(fileResult)["data"]
	fmt.Println(fileData)

	//创建待签署文件(模板)-end

	// // 创建待签署文件(本地文件)-start
	// var createFile2 Config.CreateFileConfig
	// createFile2.ContractName = "go测试合同2"
	// createFile2.ContractNo = "testContractNo02"
	// createFile2.ValidityTime = 30
	// createFile2.SignOrder = 1
	// var createFileJson2 string
	// if data, err := json.Marshal(createFile2); err == nil {
	// 	createFileJson2 = string(data)
	// }
	// var fileConfig Config.FileConfig
	// //fileConfig.FilePaths = append(fileConfig.FilePaths, "D:\\test.docx")
	// fileConfig.FileParamName = "contractFiles"
	// fileConfig.FilePaths = []string{"C:\\Users\\Administrator\\Desktop\\355.doc"}
	// Service.CreateFile(createFileJson2, &fileConfig)
	// // 创建待签署文件(本地文件)-end

	// 添加签署方-start
	var addSignerConfig Config.AddSignerConfig
	addSignerConfig.ContractNo = contractNo
	addSignerConfig.Account = "测试 Account"
	addSignerConfig.SignType = 3     // 签约方式：2：无感知签约（需要开通权限） 3：有感知签约
	addSignerConfig.ValidateType = 1 // 签署方式指定 1：短信验证码签约

	// // 定位方式 4: 模板坐标签章
	var signStrategy Config.SignStrategy
	signStrategy.AttachNo = 1
	signStrategy.LocationMode = 4
	signStrategy.SignKey = "乙方"
	signStrategy.SignPage = 7

	// 定位方式： 2坐标签章
	// var signStrategy2 Config.SignStrategy
	// signStrategy2.AttachNo = 1
	// signStrategy2.LocationMode = 2
	// signStrategy2.SignX = 0.8
	// signStrategy2.SignY = 0.8
	// signStrategy2.SignPage = 7

	// var signStrategy3 Config.SignStrategy
	// signStrategy3.AttachNo = 1
	// signStrategy3.LocationMode = 3
	// signStrategy3.SignKey = "乙方"

	addSignerConfig.SignStrategyList = append(addSignerConfig.SignStrategyList, signStrategy)

	var addSignerConfig2 Config.AddSignerConfig
	addSignerConfig2.ContractNo = contractNo
	addSignerConfig2.Account = "local_1"
	addSignerConfig2.SignType = 3
	addSignerConfig2.ValidateType = 7
	var signStrategy22 Config.SignStrategy
	signStrategy22.AttachNo = 1
	signStrategy22.LocationMode = 4
	signStrategy22.SignKey = "甲方"
	signStrategy22.SignPage = 7
	addSignerConfig2.SignStrategyList = append(addSignerConfig2.SignStrategyList, signStrategy22)
	var addSignerConfigList []Config.AddSignerConfig
	addSignerConfigList = append(addSignerConfigList, addSignerConfig, addSignerConfig2)
	var addSignerConfigJson string
	if data, err := json.Marshal(addSignerConfigList); err == nil {
		addSignerConfigJson = string(data)
	}
	Service.AddSigner(addSignerConfigJson)
}

func readFile(file string) string {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	var bodyImg io.Reader
	request, err := http.NewRequest("GET", file, bodyImg)
	if err != nil {
		log.Fatal(err)
	}

	respImg, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer respImg.Body.Close()

	srcByte, err := io.ReadAll(respImg.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := base64.StdEncoding.EncodeToString(srcByte)
	return res
}
