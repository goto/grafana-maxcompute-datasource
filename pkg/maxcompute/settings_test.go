package maxcompute

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/aliyun/aliyun-odps-go-sdk/odps"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"gotest.tools/assert"
)

func TestLoadSettings(t *testing.T) {
	t.Run("should parse settings correctly", func(t *testing.T) {
		type args struct {
			config backend.DataSourceInstanceSettings
		}
		tests := []struct {
			name         string
			args         args
			wantSettings *odps.Config
			wantErr      error
		}{
			{
				name: "should parse and set all the json fields correctly",
				args: args{
					config: backend.DataSourceInstanceSettings{
						UID:                     "ds-uid",
						JSONData:                []byte(`{ "accessKeyId": "ak", "endpoint": "endpoint", "projectName": "project", "tcpConnectionTimeout": 10, "httpTimeout": 2, "tunnelEndpoint" : "tunnelendpoint", "tunnelQuotaName": "tunnelquotaname", "others": [{"key":"aa", "value": "bb"}]}`),
						DecryptedSecureJSONData: map[string]string{"accessKeySecret": "sk", "stsToken": "sts"},
					},
				},
				wantSettings: &odps.Config{
					AccessId:             "ak",
					AccessKey:            "sk",
					StsToken:             "sts",
					Endpoint:             "endpoint",
					ProjectName:          "project",
					TcpConnectionTimeout: time.Second * 10,
					HttpTimeout:          time.Second * 2,
					TunnelEndpoint:       "tunnelendpoint",
					TunnelQuotaName:      "tunnelquotaname",
					Hints:                nil,
					Others: map[string]string{
						"aa": "bb",
					},
				},
				wantErr: nil,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotSettings, err := LoadMaxComputeConfig(tt.args.config)
				assert.Equal(t, tt.wantErr, err)
				if !reflect.DeepEqual(gotSettings, tt.wantSettings) {
					t.Errorf("LoadMaxComputeConfig() = %v, want %v", gotSettings, tt.wantSettings)
				}
			})
		}
	})
	t.Run("should capture invalid settings", func(t *testing.T) {
		tests := []struct {
			jsonData        string
			accessKeySecret string
			wantErr         error
			description     string
		}{
			{jsonData: `{ "endpoint": "" }`, accessKeySecret: "", wantErr: ErrorMessageInvalidEndpoint, description: "should capture empty endpoint"},
			{jsonData: `{ "endpoint": "foo" }`, accessKeySecret: "", wantErr: ErrorMessageInvalidProjectName, description: "should capture nil projectName"},
			{jsonData: `{ "endpoint": "foo", "projectName": "bar" }`, accessKeySecret: "", wantErr: ErrorMessageInvalidAccessKeyId, description: "should capture nil accessKeyId"},
			{jsonData: `{ "endpoint": "foo", "projectName": "bar", "accessKeyId": "baz"}`, accessKeySecret: "", wantErr: ErrorMessageInvalidAccessKeySecret, description: "should capture nil accessKeySecret"},
			{jsonData: `  "endpoint": "foo" }`, accessKeySecret: "", wantErr: ErrorMessageInvalidJSON, description: "should capture invalid json"},
		}
		for i, tc := range tests {
			t.Run(fmt.Sprintf("[%v/%v] %s", i+1, len(tests), tc.description), func(t *testing.T) {
				_, err := LoadMaxComputeConfig(backend.DataSourceInstanceSettings{
					JSONData:                []byte(tc.jsonData),
					DecryptedSecureJSONData: map[string]string{"accessKeySecret": tc.accessKeySecret},
				})
				if !errors.Is(err, tc.wantErr) {
					t.Errorf("%s not captured. %s", tc.wantErr, err.Error())
				}
			})
		}
	})
}
