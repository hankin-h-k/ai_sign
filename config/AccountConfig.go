package config

// 创建个人账号json信息配置
type CreatePersonalAccountRequest struct {
	Account         string `json:"account,omitempty"`
	SerialNo        string `json:"serialNo,omitempty"`
	Name            string `json:"name,omitempty"`
	IdCard          string `json:"idCard,omitempty"`
	IdCardType      int    `json:"idCardType,omitempty"`
	Mobile          string `json:"mobile,omitempty"`
	SignPwd         string `json:"signPwd,omitempty"`
	IsSignPwdNotice int    `json:"isSignPwdNotice,omitempty"`
	IsNotice        int    `json:"isNotice,omitempty"`
	SealPic         string `json:"sealPic,omitempty"`
}

// 基础响应
type BaseResponse struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type CreatePersonalAccountData struct {
	SealNo string `json:"sealNo,omitempty"`
}

// 创建个人账号响应
type CreatePersonalAccountResponse struct {
	BaseResponse
	Data CreatePersonalAccountData `json:"data,omitempty"`
}

// 查询用户信息json信息配置
type QueryPersonalAccountConfig struct {
	Account string `json:"account,omitempty"`
}

// 查询个人账号请求
type QueryPersonalAccountRequest struct {
	Account    string `json:"account,omitempty"`
	CreditCode string `json:"creditCode,omitempty"`
	IdCard     string `json:"idCard,omitempty"`
}

type QueryPersonalAccountData struct {
	Account      string `json:"account,omitempty"`
	Name         string `json:"name,omitempty"`
	CompanyName  string `json:"companyName,omitempty"`
	IdCard       string `json:"idCard,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	Email        string `json:"email,omitempty"`
	UserType     int    `json:"userType,omitempty"`
	CreditCode   string `json:"creditCode,omitempty"`
	BackCard     string `json:"backCard,omitempty"`
	PortVersion  int    `json:"portVersion,omitempty"`
	IdentifyType int    `json:"identifyType,omitempty"`
	AuthType     int    `json:"authType,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
	IdentifyTime string `json:"identifyTime,omitempty"`
	Status       int    `json:"status,omitempty"`
	SerialNo     string `json:"serialNo,omitempty"`
}

// 查询个人账号响应
type QueryPersonalAccountResponse struct {
	BaseResponse
	Data QueryPersonalAccountData `json:"data,omitempty"`
}

// 修改用户请求
type UpdatePersonalAccountRequest struct {
	Account      string `json:"account,omitempty"`
	Name         string `json:"name,omitempty"`
	IdentifyType int    `json:"identifyType,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	BankCard     string `json:"bankCard,omitempty"`
}

// 修改用户响应
type UpdatePersonalAccountResponse struct {
	BaseResponse
	Data string `json:"data,omitempty"`
}

// 删除用户请求
type DeletePersonalAccountRequest struct {
	Account string `json:"account,omitempty"`
}

// 删除用户响应
type DeletePersonalAccountResponse struct {
	BaseResponse
	Data string `json:"data,omitempty"`
}

// 创建企业账号json信息配置
type CreateEnterpriseAccountRequest struct {
	Account         string `json:"account,omitempty"`
	SerialNo        string `json:"serialNo,omitempty"`
	CompanyName     string `json:"companyName,omitempty"`
	CreditCode      string `json:"creditCode,omitempty"`
	CreditType      int    `json:"creditType,omitempty"`
	Name            string `json:"name,omitempty"`
	IdCard          string `json:"idCard,omitempty"`
	IdCardType      int    `json:"idCardType,omitempty"`
	Mobile          string `json:"mobile,omitempty"`
	ContactName     string `json:"contactName,omitempty"`
	ContactIdCard   string `json:"contactIdCard,omitempty"`
	SignPwd         string `json:"signPwd,omitempty"`
	IsSignPwdNotice int    `json:"isSignPwdNotice,omitempty"`
	IsNotice        int    `json:"isNotice,omitempty"`
}

// 创建企业账号响应数据
type CreateEnterpriseAccountData struct {
	SealNo string `json:"sealNo,omitempty"`
}

// 创建企业账号响应
type CreateEnterpriseAccountResponse struct {
	BaseResponse
	Data CreateEnterpriseAccountData `json:"data,omitempty"`
}

// 查询企业用户请求参数
type QueryEnterpriseAccountRequest struct {
	Account     string `json:"account,omitempty"`
	CreditCode  string `json:"creditCode,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
}

type QueryEnterpriseAccountData struct {
	Account      string `json:"account,omitempty"`
	Name         string `json:"name,omitempty"`
	CompanyName  string `json:"companyName,omitempty"`
	IdCard       string `json:"idCard,omitempty"`
	UserType     int    `json:"userType,omitempty"`
	Status       int    `json:"status,omitempty"`
	CreditCode   string `json:"creditCode,omitempty"`
	BackCard     string `json:"backCard,omitempty"`
	PortVersion  int    `json:"portVersion,omitempty"`
	IdentifyType int    `json:"identifyType,omitempty"`
	AuthType     int    `json:"authType,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
	IdentifyTime string `json:"identifyTime,omitempty"`
	SerialNo     string `json:"serialNo,omitempty"`
}

// 查询企业用户响应
type QueryEnterpriseAccountResponse struct {
	BaseResponse
	Data []QueryEnterpriseAccountData `json:"data,omitempty"`
}

// 修改企业用户请求
type UpdateEnterpriseAccountRequest struct {
	Account     string `json:"account,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	Name        string `json:"name,omitempty"`
	IdCard      string `json:"idCard,omitempty"`
	IdCardType  int    `json:"idCardType,omitempty"`
}

// 修改企业用户响应
type UpdateEnterpriseAccountResponse struct {
	BaseResponse
	Data string `json:"data,omitempty"`
}
