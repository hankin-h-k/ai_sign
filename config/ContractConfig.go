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
	FileName      string                   `json:"fileName,omitempty"`
	ContractNo    string                   `json:"contractNo,omitempty"`
	FillData      map[string]string        `json:"fillData,omitempty"`
	ComponentData []map[string]interface{} `json:"componentData,omitempty"`
	TableDatas    []map[string]interface{} `json:"tableDatas,omitempty"`
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
	SealNo       int     `json:"sealNo,omitempty"`
	CanDrag      int     `json:"canDrag,omitempty"`
	SignKey      string  `json:"signKey,omitempty"`
	SignType     int     `json:"signType,omitempty"`
	SignPage     int     `json:"signPage,omitempty"`
	SignX        float64 `json:"signX,omitempty"`
	SignY        float64 `json:"signY,omitempty"`
	OffsetX      float64 `json:"offsetX,omitempty"`
	OffsetY      float64 `json:"offsetY,omitempty"`
}

type SignStrike struct {
	AttachNo   int     `json:"attachNo,omitempty"`
	CrossStyle int     `json:"crossStyle,omitempty"`
	SignPage   int     `json:"signPage,omitempty"`
	OddEven    int     `json:"oddEven,omitempty"`
	OffsetX    float64 `json:"offsetX,omitempty"`
	OffsetY    float64 `json:"offsetY,omitempty"`
	SignKey    string  `json:"signKey,omitempty"`
	ClipNumber int     `json:"clipNumber,omitempty"`
}

type Options struct {
	Index    int  `json:"index,omitempty"`
	Selected bool `json:"selected,omitempty"`
}

type RowValues struct {
	InsertRow bool     `json:"insertRow,omitempty"`
	ColValues []string `json:"colValues,omitempty"`
}

type ReceiverFillStrategy struct {
	AttachNo  int         `json:"attachNo,omitempty"`
	Key       string      `json:"key,omitempty"`
	Value     string      `json:"value,omitempty"`
	FillStage int         `json:"fillStage,omitempty"`
	Options   []Options   `json:"options,omitempty"`
	RowValues []RowValues `json:"rowValues,omitempty"`
}

// 下载合同
type DownloadFileConfig struct {
	ContractNo string `json:"contractNo,omitempty"`
}

// 创建待签署文件请求
type CreateContractRequest struct {
	ContractNo           string      `json:"contractNo,omitempty"`
	ContractName         string      `json:"contractName,omitempty"`
	ValidityTime         int         `json:"validityTime,omitempty"`
	ValidityDate         string      `json:"validityDate,omitempty"`
	SignOrder            int         `json:"signOrder,omitempty"`
	FileIds              []string    `json:"fileIds,omitempty"`
	ContractFiles        [][]byte    `json:"contractFiles,omitempty"`
	Templates            []Templates `json:"templates,omitempty"`
	ReadSeconds          int         `json:"readSeconds,omitempty"`
	ReadType             int         `json:"readType,omitempty"`
	NeedAgree            int         `json:"needAgree,omitempty"`
	AutoExpand           int         `json:"autoExpand,omitempty"`
	NotifyUrl            string      `json:"notifyUrl,omitempty"`
	CallbackUrl          string      `json:"callbackUrl,omitempty"`
	UserNotifyUrl        string      `json:"userNotifyUrl,omitempty"`
	RedirectUrl          string      `json:"redirectUrl,omitempty"`
	RefuseOn             int         `json:"refuseOn,omitempty"`
	AutoContinue         int         `json:"autoContinue,omitempty"`
	ViewFlg              int         `json:"viewFlg,omitempty"`
	RedirectReturnUrl    string      `json:"RedirectReturnUrl,omitempty"`
	RedirectCompletedUrl string      `json:"redirectCompletedUrl,omitempty"`
}

type ContractFiles struct {
	FileName string `json:"fileName,omitempty"`
	AttachNo int    `json:"attachNo,omitempty"`
	Page     int    `json:"page,omitempty"`
}

type CreateContractData struct {
	PreviewUrl    string          `json:"previewUrl,omitempty"`
	ContractFiles []ContractFiles `json:"contactFiles,omitempty"`
}

// 创建待签署文件响应
type CreateContractResponse struct {
	BaseResponse
	Data CreateContractData `json:"data,omitempty"`
}

type AddSignerRequest struct {
	ContractNo               string                 `json:"contractNo,omitempty"`
	Account                  string                 `json:"account,omitempty"`
	SignType                 int                    `json:"signType,omitempty"`
	SealNo                   string                 `json:"sealNo,omitempty"`
	AuthSignAccount          string                 `json:"authSignAccount,omitempty"`
	NoticeMobile             string                 `json:"noticeMobile,omitempty"`
	SignOrder                string                 `json:"signOrder,omitempty"`
	IsNotice                 string                 `json:"is_notice,omitempty"`
	ValidateType             int                    `json:"validateType,omitempty"`
	FaceAuthMode             int                    `json:"faceAuthMode,omitempty"`
	ValidateTypeList         string                 `json:"validateTypeList,omitempty"`
	AutoSwitch               int                    `json:"autoSwitch,omitempty"`
	IsNoticeComplete         int                    `json:"isNoticeComplete,omitempty"`
	WaterMark                int                    `json:"waterMark,omitempty"`
	AutoSms                  int                    `json:"autoSms,omitempty"`
	CustomSignFlag           int                    `json:"customSignFlag,omitempty"`
	SignStrategyList         []SignStrategy         `json:"signStrategyList,omitempty"`
	SignStrikeList           []SignStrike           `json:"signStrikeList,omitempty"`
	ReceiverFillStrategyList []ReceiverFillStrategy `json:"receiverFillStrategyList,omitempty"`
	AuthConfig               map[string]interface{} `json:"authConfig,omitempty"`
	IsIframe                 int                    `json:"isIframe,omitempty"`
	WillType                 string                 `json:"willType,omitempty"`
	SignMark                 string                 `json:"signMark,omitempty"`
}

type SignUser struct {
	Account    string `json:"account,omitempty"`
	SignUrl    string `json:"signUrl,omitempty"`
	PwdSignUrl string `json:"pwdSignUrl,omitempty"`
	SignOrder  int    `json:"signOrder,omitempty"`
	Name       string `json:"name,omitempty"`
	IdCard     string `json:"idCard,omitempty"`
	SignMark   string `json:"signMark,omitempty"`
}

type AddSignerData struct {
	ContractNo   string     `json:"contractNo,omitempty"`
	ContractName string     `json:"contractName,omitempty"`
	ValidityTime int        `json:"validityTime,omitempty"`
	PreviewUrl   string     `json:"previewUrl,omitempty"`
	SignUser     []SignUser `json:"signUser,omitempty"`
}

type AddSignerResponse struct {
	BaseResponse
	Data AddSignerData `json:"data,omitempty"`
}

type QueryContractStatusRequest struct {
	ContractNo string `json:"contractNo,omitempty"`
}

type QueryContractStatusData struct {
	ContractNo   string `json:"contractNo,omitempty"`
	ContractName string `json:"contractName,omitempty"`
	Status       int    `json:"status,omitempty"`
}

type QueryContractStatusResponse struct {
	BaseResponse
	Data QueryContractStatusData `json:"data,omitempty"`
}

type QueryContractRequest struct {
	ContractNo string `json:"contractNo,omitempty"`
}

type ContractViewHis struct {
	AccessTime string `json:"accessTime,omitempty"`
	Detail     string `json:"detail,omitempty"`
	Opened     int    `json:"opened,omitempty"`
}

type ContractSignUser struct {
	Account          string            `json:"account,omitempty"`
	SignUrl          string            `json:"signUrl,omitempty"`
	PwdSignUrl       string            `json:"pwdSignUrl,omitempty"`
	SignOrder        int               `json:"signOrder,omitempty"`
	Name             string            `json:"name,omitempty"`
	IdCard           string            `json:"idCard,omitempty"`
	Mobile           string            `json:"mobile,omitempty"`
	CompanyName      string            `json:"companyName,omitempty"`
	UserType         int               `json:"userType,omitempty"`
	SealNo           string            `json:"sealNo,omitempty"`
	SignType         string            `json:"signType,omitempty"`
	ValidateType     int               `json:"validateType,omitempty"`
	SignStatus       int               `json:"signStatus,omitempty"`
	SignFinishedTime string            `json:"signFinishedTime,omitempty"`
	SignMark         string            `json:"signMark,omitempty"`
	ContractViewHis  []ContractViewHis `json:"contractViewHis,omitempty"`
}

type QueryContractData struct {
	ContractNo   string             `json:"contractNo,omitempty"`
	Status       int                `json:"status,omitempty"`
	ContractName string             `json:"contractName,omitempty"`
	ValidityTime int                `json:"validityTime,omitempty"`
	PreviewUrl   string             `json:"previewUrl,omitempty"`
	EmbeddedUrl  string             `json:"embeddedUrl"`
	Remark       string             `json:"remark,omitempty"`
	SignUser     []ContractSignUser `json:"signUser,omitempty"`
}

type QueryContractResponse struct {
	BaseResponse
	Data QueryContractData `json:"data,omitempty"`
}

type PreviewContractRequest struct {
	ContractNo       string         `json:"contractNo,omitempty"`
	Account          string         `json:"account,omitempty"`
	IsWrite          int            `json:"isWrite,omitempty"`
	SignStrategyList []SignStrategy `json:"signStrategyList,omitempty"`
}

type PreviewContractResponse struct {
	BaseResponse
	Data string `json:"data,omitempty"`
}

type DownloadContractRequest struct {
	ContractNo       string `json:"contractNo,omitempty"`
	Outfile          string `json:"outfile,omitempty"`
	Force            int    `json:"force,omitempty"`
	DownloadFileType int    `json:"downloadFileType,omitempty"`
}

type DownloadContractData struct {
	FileName string `json:"filename,omitempty"`
	Md5      string `json:"md5,omitempty"`
	FileType int    `json:"fileType,omitempty"`
	Size     int64  `json:"size,omitempty"`
	Data     string `json:"data,omitempty"`
}

type DownloadContractRespnse struct {
	BaseResponse
	Data DownloadContractData `json:"data"`
}

type DownloadContractByAttachRequest struct {
	ContractNo       string   `json:"contractNo,omitempty"`
	AttachNames      []string `json:"attachNames,omitempty"`
	Outfile          string   `json:"outfile,omitempty"`
	DownloadFileType int      `json:"downloadFileType,omitempty"`
}

type DownloadContractByAttachRespnse struct {
	BaseResponse
	Data DownloadContractData `json:"data"`
}

type WithdrawContractRequest struct {
	ContractNo       string `json:"contractNo,omitempty"`
	WithdrawReason   string `json:"withdrawReason,omitempty"`
	IsNoticeSignUser bool   `json:"isNoticeSignUser,omitempty"`
}

type WithdrawContractResponse struct {
	BaseResponse
}

type CancellationContractRequest struct {
	ContractNo   string `json:"contractNo,omitempty"`
	ValidityTime int    `json:"validityTime,omitempty"`
}

type CancellationContractData struct {
	ContractNo   string   `json:"contractNo,omitempty"`
	ContractName string   `json:"contractName,omitempty"`
	ValidityTime int      `json:"validityTime,omitempty"`
	SignUser     SignUser `json:"signUser,omitempty"`
}

type CancellationContractResponse struct {
	BaseResponse
	Data CancellationContractData `json:"data"`
}

type DelayContractRequest struct {
	ContractNo   string `json:"contractNo,omitempty"`
	ValidityDate string `json:"validityDate,omitempty"`
}

type DelayContractResponse struct {
	BaseResponse
}
