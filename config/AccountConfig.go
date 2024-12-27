package config



//创建个人账号json信息配置
type CreatePersonalAccountConfig struct {
	Account string `json:"account,omitempty"`
	SerialNo string `json:"serialNo,omitempty"`
	Name string `json:"name,omitempty"`
	IdCard string `json:"idCard,omitempty"`
	IdCardType int `json:"idCardType,omitempty"`
	Mobile string `json:"mobile,omitempty"`

}

// 查询用户信息json信息配置
type QueryPersonalAccountConfig struct {
	Account string `json:"account,omitempty"`
}
