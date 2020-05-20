package sms

import (
	"log"
	"testing"
)

const (
	accessKeyID     = ""
	accessKeySecret = ""
	signName        = ""
)

func TestSendSms(t *testing.T) {
	if err := NewClient(accessKeyID, accessKeySecret, signName); err != nil {
		log.Fatal(err)
	}
	templateParamMap := make(map[string]interface{})
	templateParamMap["name"] = "王松"
	templateParamMap["drop_rate"] = 50
	templateParamMap["organization"] = "江苏卫川物联科技有限公司"
	smsResponse, err := Send("13205173164", "SMS_190272170", templateParamMap)

	if err != nil {
		log.Fatal(err)
	}

	if smsResponse.Code != "OK" {
		log.Fatal(smsResponse.Message)
	}
}
