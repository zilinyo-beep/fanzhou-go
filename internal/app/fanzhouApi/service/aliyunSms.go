package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type AliyunSmsService struct {
	client *dysmsapi.Client
	cfg    AliyunSmsConfig
}

type AliyunSmsConfig struct {
	AccessKeyId     string
	AccessKeySecret string
	TemplateCode    string
}

var (
	smsService     *AliyunSmsService
	smsServiceOnce sync.Once
)

func GetAliyunSmsService() *AliyunSmsService {
	smsServiceOnce.Do(func() {
		smsService, _ = NewAliyunSmsService()
	})
	return smsService
}

func NewAliyunSmsService() (*AliyunSmsService, error) {
	ctx := gctx.New()
	cfg := AliyunSmsConfig{
		AccessKeyId:     g.Cfg().MustGet(ctx, "aliyun.sms.accessKeyId").String(),
		AccessKeySecret: g.Cfg().MustGet(ctx, "aliyun.sms.accessKeySecret").String(),
		TemplateCode:    g.Cfg().MustGet(ctx, "aliyun.sms.templateCode").String(),
	}

	client, err := dysmsapi.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(cfg.AccessKeyId),
		AccessKeySecret: tea.String(cfg.AccessKeySecret),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	})
	if err != nil {
		return nil, fmt.Errorf("创建短信 client 失败: %w", err)
	}

	return &AliyunSmsService{
		client: client,
		cfg:    cfg,
	}, nil
}

func (s *AliyunSmsService) SendSMS(phone, signName string, params map[string]string) error {
	paramJson, _ := json.Marshal(params)

	req := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(signName),
		TemplateCode:  tea.String(s.cfg.TemplateCode),
		TemplateParam: tea.String(string(paramJson)),
	}

	resp, err := s.client.SendSmsWithOptions(req, &util.RuntimeOptions{})
	if err != nil {
		return fmt.Errorf("短信发送失败: %w", err)
	}

	var body map[string]interface{}
	_ = json.Unmarshal([]byte(resp.Body.String()), &body)
	if body["Code"] != "OK" {
		return errors.New(fmt.Sprintf("短信失败：%v", body["Message"]))
	}

	return nil
}
