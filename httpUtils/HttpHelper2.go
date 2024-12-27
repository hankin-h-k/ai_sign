package httpUtils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"time"

	Config "github.com/hankin-h-k/ai_sign/config"
)

func SendRequest2(urlStr, jsonStr string, fileConfig *Config.FileConfig) ([]byte, error) {
	// 生成随机边界
	boundary := randomBoundary()

	var appId = Config.GetInitConfig().AppId()
	var privateKey = Config.GetInitConfig().PrivateKey()
	var host = Config.GetInitConfig().Host() + urlStr

	// 获取当前时间戳并加10分钟
	timestamp := time.Now().Add(10*time.Minute).UnixNano() / int64(time.Millisecond)

	// 对 bizData 进行排序
	sortedJsonStr, err := sortJSON2(jsonStr)
	if err != nil {
		log.Println("排序 JSON 失败:", err)
		return nil, err
	}

	log.Println("排序后的JSON字符串:", sortedJsonStr)

	// 生成签名
	rsaSuffix := sortedJsonStr + md5Hex(sortedJsonStr) + appId + strconv.FormatInt(timestamp, 10)

	sign, err := generateSHA1withRSASignature(rsaSuffix, privateKey)
	if err != nil {
		log.Println("生成签名失败:", err)
		return nil, err
	}
	// 设置请求头部
	headers := map[string]string{
		"sign":      sign,
		"timestamp": strconv.FormatInt(timestamp, 10),
	}

	// 创建 HTTP 客户端
	var proxy *url.URL
	if proxyHost := os.Getenv("PROXY_HOST"); proxyHost != "" {
		proxyPort, _ := strconv.Atoi(os.Getenv("PROXY_PORT"))
		proxy = &url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("%s:%d", proxyHost, proxyPort),
		}
	}

	client := &http.Client{
		Timeout: TIME_OUT,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}

	// 创建请求
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加请求体参数
	writer.WriteField("appId", appId)
	writer.WriteField("timestamp", strconv.FormatInt(timestamp, 10))
	writer.WriteField("bizData", jsonStr)
	writer.SetBoundary(boundary)

	// 添加文件
	if fileConfig != nil {
		for _, filePath := range fileConfig.FilePaths {
			file, err := os.Open(filePath)
			if err != nil {
				return nil, err
			}
			defer file.Close()

			part, err := writer.CreateFormFile(fileConfig.FileParamName, filePath)
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(part, file)
			if err != nil {
				return nil, err
			}
		}
	}
	// 关闭 writer
	writer.Close()
	// 创建请求
	req, err := http.NewRequest("POST", host, body)
	if err != nil {
		log.Println("创建请求失败:", err)
		return nil, err
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Charset", "UTF-8")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("发送请求失败:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应失败:", err)
		return nil, err
	}
	log.Println("调用结果:", string(respBody))
	return respBody, nil
}

// 对 map 的 key 进行排序
func sortMapKeys(data map[string]interface{}) map[string]interface{} {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sortedData := make(map[string]interface{})
	for _, k := range keys {
		sortedData[k] = data[k]
	}
	return sortedData
}

// 递归排序 JSON is arrays
func sortJSON2(jsonStr string) (string, error) {
	var data []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return "", err
	}
	for i, item := range data {
		sortedItem := sortMapKeys(item)
		data[i] = sortedItem
	}
	sortedJson, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(sortedJson), nil
}
