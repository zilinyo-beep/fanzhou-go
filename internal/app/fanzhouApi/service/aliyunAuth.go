package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dypnsapi "github.com/alibabacloud-go/dypnsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type AliyunAuthConfig struct {
	AccessKeyId     string
	AccessKeySecret string
	BundleId        string
	DurationSeconds int64
}

type AuthResult struct {
	Code               string            `json:"code"`
	Message            string            `json:"message"`
	GetMobileResultDTO map[string]string `json:"getMobileResultDTO"`
	RequestId          string            `json:"requestId"`
}

type AliyunAuthService struct {
	client *dypnsapi.Client
	cfg    AliyunAuthConfig
}

var (
	instance *AliyunAuthService
	once     sync.Once
)

func GetAliyunAuthService() (*AliyunAuthService, error) {
	var initErr error
	once.Do(func() {
		ctx := gctx.New()
		cfg := AliyunAuthConfig{
			AccessKeyId:     g.Cfg().MustGet(ctx, "aliyun.ident.accessKeyId").String(),
			AccessKeySecret: g.Cfg().MustGet(ctx, "aliyun.ident.accessKeySecret").String(),
			BundleId:        g.Cfg().MustGet(ctx, "aliyun.ident.bundleId").String(),
			DurationSeconds: g.Cfg().MustGet(ctx, "aliyun.ident.durationSeconds").Int64(),
		}

		client, err := dypnsapi.NewClient(&openapi.Config{
			AccessKeyId:     tea.String(cfg.AccessKeyId),
			AccessKeySecret: tea.String(cfg.AccessKeySecret),
			Endpoint:        tea.String("dypnsapi.aliyuncs.com"),
		})
		if err != nil {
			initErr = fmt.Errorf("创建一键登录 client 失败: %w", err)
			return
		}

		instance = &AliyunAuthService{
			client: client,
			cfg:    cfg,
		}
	})

	return instance, initErr
}

// GetPhoneByToken 根据 token 获取手机号
func (s *AliyunAuthService) GetPhoneByToken(token string) (string, error) {
	req := &dypnsapi.GetMobileRequest{
		AccessToken: tea.String(token),
		OutId:       tea.String(s.cfg.BundleId),
	}

	resp, err := s.client.GetMobileWithOptions(req, &util.RuntimeOptions{})
	if err != nil {
		return "", fmt.Errorf("一键登录请求失败: %w", err)
	}
	if resp == nil || resp.Body == nil {
		return "", errors.New("响应为空")
	}

	var body AuthResult
	err = json.Unmarshal([]byte(resp.Body.String()), &body)
	if err != nil {
		return "", err
	}
	mobile := body.GetMobileResultDTO["Mobile"]
	if mobile == "" {
		return "", errors.New("未获取到手机号")
	}
	return mobile, nil
}
