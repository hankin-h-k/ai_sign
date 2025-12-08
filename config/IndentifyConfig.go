package config

type PersonIdentifyUrlRequest struct {
	AuthTypeList []int `json:"authTypeList"`
	RealName string `json:"realName"`
	IdCardNo string `json:"idCardNo"`
	Mobile string `json:"mobile"`
	BankCard string `json:"bankCard"`
	RedirectUrl string `json:"redirectUrl"`
	NotifyUrl string `json:"notifyUrl"`
	BizId string `json:"bizId"`
	FaceAuthMode int `json:"faceAuthMode"`
	Resubmit int `json:"resubmit"`
	ShowResult int `json:"showResult"`
	MobileRequired int `json:"mobileRequired"`
	IsIframe int `json:"isIframe"`
	IsNotice int `json:"isNotice"`
	NoticeMobile string `json:"noticeMobile"`
	NeedSignPwd int  `json:"needSignPwd"`
}

type PersonIdentifyUrlData struct {
	Result int `json:"result"`
	SerialNo string `json:"serialNo"`
	Type string `json:"type"`
	IdentifyUrl string `json:"identifyUrl"`
}

type PersonIdentifyUrlResponse struct {
	BaseResponse
	Data PersonIdentifyUrlData `json:"data"`
}

type AuthRecordInfoRequest struct {
	SerialNo string `json:"serialNo"`
}

type AuthRecordInfoData struct {
	SerialNo string `json:"serialNo"`
	UserType int `json:"userType"`
	Result int `json:"result"`
	RealName string `json:"realName"`
	IdCardNo string `json:"idCardNo"`
	Mobile string `json:"mobile"`
	CompanyName string `json:"companyName"`
	CreditCode string `json:"creditCode"`
	AuthTypeCode string `json:"authTypeCode"`
	AuthTypeName string `json:"authTypeName"`
	AgentResult int `json:"agentResult"`
	AgentName string `json:"agentName"`
	AgentIdCardNo string `json:"agentIdCardNo"`
	AgentMobile string `json:"agentMobile"`
	AgentAuthTypeCode	int `json:"agentAuthTypeCode"`
	AgentAuthTypeName string `json:"agentAuthTypeName"`
}

type AuthRecordInfoResponse struct {
	BaseResponse
	Data AuthRecordInfoData `json:"data"`
}