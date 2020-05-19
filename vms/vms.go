package vms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dyvmsapi"
)

type Response struct {
	RequestId string `json:"RequestId" xml:"RequestId"`
	CallId    string `json:"CallId" xml:"CallId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

var (
	vmsClient        *dyvmsapi.Client
	calledShowNumber string
)

func NewClient(accessKeyId, accessKeySecret, calledNumber string) error {
	var err error
	calledShowNumber = calledNumber
	vmsClient, err = dyvmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessKeySecret)
	return err
}

func Call(calledNumber, ttsCode, ttsParam string) (*Response, error) {
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
	return &Response{
		RequestId: response.RequestId,
		Code:      response.Code,
		Message:   response.Message,
		CallId:    response.CallId,
	}, nil
}
