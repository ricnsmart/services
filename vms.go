package services

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dyvmsapi"
)

type VmsResponse struct {
	RequestId string `json:"RequestId" xml:"RequestId"`
	CallId    string `json:"CallId" xml:"CallId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

var vmsClient *dyvmsapi.Client

func InitAliVms(accessKeyId, accessKeySecret string) error {
	var err error
	vmsClient, err = dyvmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessKeySecret)
	return err
}

func Call(calledNumber, calledShowNumber, ttsCode, ttsParam string) (*VmsResponse, error) {
	request := dyvmsapi.CreateSingleCallByTtsRequest()
	request.Scheme = "https"
	request.CalledNumber = calledNumber
	request.CalledShowNumber = calledShowNumber
	request.TtsCode = ttsCode
	request.TtsParam = ttsParam
	response, err := vmsClient.SingleCallByTts(request)
	if err != nil {
		return nil, err
	}
	return &VmsResponse{
		RequestId: response.RequestId,
		Code:      response.Code,
		Message:   response.Message,
		CallId:    response.CallId,
	}, nil
}
