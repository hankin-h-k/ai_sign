package config

type InitConfig struct {
	host string
	appId string
	privateKey string
}

var initConfig *InitConfig = new(InitConfig)

 func GetInitConfig() *InitConfig {
	return initConfig
}

func (e *InitConfig) Host() string {
	return e.host
}

func (e *InitConfig) SetHost(host string)  {
	e.host = host
}

func (e *InitConfig) AppId() string  {
	return e.appId
}

func (e *InitConfig) SetAppId(appId string)  {
	e.appId = appId
}

func (e *InitConfig) PrivateKey() string  {
	return e.privateKey
}

func (e *InitConfig) SetPrivateKey(privateKey string)  {
	e.privateKey = privateKey
}
