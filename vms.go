package services

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dyvmsapi"
	"log"
)

type VmsResponse struct {
	RequestId string `json:"RequestId" xml:"RequestId"`
	CallId    string `json:"CallId" xml:"CallId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

var vmsClient *dyvmsapi.Client

func InitVms(accessKeyId, accessKeySecret string) {
	var err error
	vmsClient, err = dyvmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessKeySecret)
	if err != nil {
		log.Fatal(err)
	}
}

func Call(calledNumber, ttsCode, ttsParam string) (*VmsResponse, error) {
	request := dyvmsapi.CreateSingleCallByTtsRequest()
	request.Scheme = "https"
	request.CalledNumber = calledNumber
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
