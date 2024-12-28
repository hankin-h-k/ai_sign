package config

type FileConfig struct {
	FileParamName string
	FilePaths     []string
}

// 查询文件状态请求
type QueryFileInfoRequest struct {
	FileId string `json:"fileId,omitempty"`
}

type QueryFileInfoData struct {
	FileId      string `json:"fileId,omitempty"`
	FileName    string `json:"filename,omitempty"`
	Extension   string `json:"extension,omitempty"`
	Filesize    int64  `json:"filesize,omitempty"`
	ContentMd5  string `json:"contentMd5,omitempty"`
	IsUpload    int    `json:"isUpload,omitempty"`
	DownloadUrl string `json:"downloadUrl,omitempty"`
}

type QueryFileInfoResponse struct {
	BaseResponse
	Data QueryFileInfoData `json:"data,omitempty"`
}
