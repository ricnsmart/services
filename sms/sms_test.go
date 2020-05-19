package sms

import (
	"log"
	"testing"
)

const (
	accessKeyID     = "LTAI5Gw0wvWgF0QD"
	accessKeySecret = "TB2lpHjV5QMcF7zRHG2Ct6RnkWffnN"
	signName        = "卫川智联"
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
