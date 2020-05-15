package services

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type SmsResponse struct {
	RequestId string `json:"RequestId" xml:"RequestId"`
	BizId     string `json:"BizId" xml:"BizId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

var (
	smsClient   *dysmsapi.Client
	smsSignName string
)

func InitAliSms(accessKeyId, accessKeySecret, signName string) error {
	var err error
	smsSignName = signName
	smsClient, err = dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessKeySecret)
	return err
}

func SendSms(phoneNumbers, templateCode, templateParam string) (*SmsResponse, error) {
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = smsSignName
	request.PhoneNumbers = phoneNumbers
	request.TemplateCode = templateCode
	request.TemplateParam = templateParam
	response, err := smsClient.SendSms(request)
	if err != nil {
		return nil, err
	}
	return &SmsResponse{
		RequestId: response.RequestId,
		Code:      response.Code,
		Message:   response.Message,
		BizId:     response.BizId,
	}, nil
}
