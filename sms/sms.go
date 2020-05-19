package sms

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"strings"
)

type Response struct {
	RequestId string `json:"RequestId" xml:"RequestId"`
	BizId     string `json:"BizId" xml:"BizId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

var (
	Client      *dysmsapi.Client
	smsSignName string
)

func NewClient(accessKeyId, accessKeySecret, signName string) error {
	var err error
	smsSignName = signName
	Client, err = dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessKeySecret)
	return err
}

func Send(phoneNumbers, templateCode string, templateParamMap map[string]interface{}) (*Response, error) {
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = smsSignName
	request.PhoneNumbers = phoneNumbers
	request.TemplateCode = templateCode

	var templateParamArr []string
	for key, value := range templateParamMap {
		templateParamArr = append(templateParamArr, fmt.Sprintf(`"%v":"%v"`, key, value))
	}

	request.TemplateParam = fmt.Sprintf(`{%v}`, strings.Join(templateParamArr, ","))

	response, err := Client.SendSms(request)
	if err != nil {
		return nil, err
	}
	return &Response{
		RequestId: response.RequestId,
		Code:      response.Code,
		Message:   response.Message,
		BizId:     response.BizId,
	}, nil
}
