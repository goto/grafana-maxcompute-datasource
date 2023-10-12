package maxcompute

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aliyun/aliyun-odps-go-sdk/odps"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

type CustomOption struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func isValid(config *odps.Config) (err error) {
	if config.Endpoint == "" {
		return ErrorMessageInvalidEndpoint
	}

	if config.ProjectName == "" {
		return ErrorMessageInvalidProjectName
	}

	if config.AccessId == "" {
		return ErrorMessageInvalidAccessKeyId
	}

	if config.AccessKey == "" {
		return ErrorMessageInvalidAccessKeySecret
	}

	return nil
}

func LoadMaxComputeConfig(settings backend.DataSourceInstanceSettings) (*odps.Config, error) {

	var jsonData map[string]interface{}
	if err := json.Unmarshal(settings.JSONData, &jsonData); err != nil {
		return nil, fmt.Errorf("%s: %w", err.Error(), ErrorMessageInvalidJSON)
	}

	config := odps.NewConfig()

	if jsonData["endpoint"] != nil {
		config.Endpoint = jsonData["endpoint"].(string)
	}

	if jsonData["projectName"] != nil {
		config.ProjectName = jsonData["projectName"].(string)
	}

	if jsonData["accessKeyId"] != nil {
		config.AccessId = jsonData["accessKeyId"].(string)
	}

	if jsonData["tcpConnectionTimeout"] != nil {
		timeout := int64(jsonData["tcpConnectionTimeout"].(float64))
		config.TcpConnectionTimeout = time.Second * time.Duration(timeout)
	}

	if jsonData["httpTimeout"] != nil {
		timeout := int64(jsonData["httpTimeout"].(float64))
		config.HttpTimeout = time.Second * time.Duration(timeout)
	}

	if jsonData["tunnelEndpoint"] != nil {
		config.TunnelEndpoint = jsonData["tunnelEndpoint"].(string)
	}

	if jsonData["tunnelQuotaName"] != nil {
		config.TunnelQuotaName = jsonData["tunnelQuotaName"].(string)
	}

	if jsonData["others"] != nil {
		if optionBytes, err := json.Marshal(jsonData["others"]); err == nil {
			options := make([]CustomOption, 0)
			err = json.Unmarshal(optionBytes, &options)
			if err == nil {
				config.Others = make(map[string]string)
				for _, v := range options {
					config.Others[v.Key] = v.Value
				}
			}
		}
	}

	if accessSecret, ok := settings.DecryptedSecureJSONData["accessKeySecret"]; ok {
		config.AccessKey = accessSecret
	}

	if stsToken, ok := settings.DecryptedSecureJSONData["stsToken"]; ok {
		config.StsToken = stsToken
	}

	return config, isValid(config)
}
