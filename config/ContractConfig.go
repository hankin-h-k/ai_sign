package config

// 创建待签署文件
type CreateFileConfig struct {
	ContractNo   string      `json:"contractNo,omitempty"`
	ContractName string      `json:"contractName,omitempty"`
	ValidityTime int         `json:"validityTime,omitempty"`
	SignOrder    int         `json:"signOrder,omitempty"`
	Templates    []Templates `json:"templates,omitempty"`
}

type Templates struct {
	TemplateNo    string                   `json:"templateNo,omitempty"`
	FillData      map[string]string        `json:"fillData,omitempty"`
	ComponentData []map[string]interface{} `json:"componentData,omitempty"`
}

// 添加签署方
type AddSignerConfig struct {
	ContractNo       string         `json:"contractNo,omitempty"`
	Account          string         `json:"account,omitempty"`
	SignType         int            `json:"signType,omitempty"`
	ValidateType     int            `json:"validateType,omitempty"`
	SignStrategyList []SignStrategy `json:"signStrategyList,omitempty"`
}

type SignStrategy struct {
	AttachNo     int     `json:"attachNo,omitempty"`
	LocationMode int     `json:"locationMode,omitempty"`
	SignKey      string  `json:"signKey,omitempty"`
	SignPage     int     `json:"signPage,omitempty"`
	SignX        float64 `json:"signX,omitempty"`
	SignY        float64 `json:"signY,omitempty"`
}

// 下载合同
type DownloadFileConfig struct {
	ContractNo string `json:"contractNo,omitempty"`
}
