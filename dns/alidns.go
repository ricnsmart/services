package dns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type Record struct {
	Value      string `json:"Value" xml:"Value"`
	TTL        int64  `json:"TTL" xml:"TTL"`
	Remark     string `json:"Remark" xml:"Remark"`
	DomainName string `json:"DomainName" xml:"DomainName"`
	RR         string `json:"RR" xml:"RR"`
	Priority   int64  `json:"Priority" xml:"Priority"`
	RecordId   string `json:"RecordId" xml:"RecordId"`
	Status     string `json:"Status" xml:"Status"`
	Locked     bool   `json:"Locked" xml:"Locked"`
	Weight     int    `json:"Weight" xml:"Weight"`
	Line       string `json:"Line" xml:"Line"`
	Type       string `json:"Type" xml:"Type"`
}

type GetDomainRecordsResponse struct {
	RequestId     string   `json:"RequestId" xml:"RequestId"`
	Total         int64    `json:"Total" xml:"Total"`
	DomainRecords []Record `json:"DomainRecords" xml:"DomainRecords"`
}

type DomainRecordResponse struct {
	RequestId string `json:"RequestId" xml:"RequestId"`
	RecordId  string `json:"RecordId" xml:"RecordId"`
}

var Client *alidns.Client

func NewClient(accessKeyId, accessKeySecret string) error {
	var err error
	Client, err = alidns.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessKeySecret)
	return err
}

func GetDomainRecords(domainName, keyWord string, pageNumber, pageSize int) (*GetDomainRecordsResponse, error) {
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.Scheme = "https"
	request.DomainName = domainName
	request.PageNumber = requests.NewInteger(pageNumber)
	request.PageSize = requests.NewInteger(pageSize)
	request.KeyWord = keyWord

	response, err := Client.DescribeDomainRecords(request)
	if err != nil {
		return nil, err
	}

	var domainRecords []Record
	for _, r := range response.DomainRecords.Record {
		var record Record
		record.Value = r.Value
		record.TTL = r.TTL
		record.Remark = r.Remark
		record.DomainName = r.DomainName
		record.RR = r.RR
		record.Priority = r.Priority
		record.RecordId = r.RecordId
		record.Status = r.Status
		record.Locked = r.Locked
		record.Weight = r.Weight
		record.Line = r.Line
		record.Type = r.Type
		domainRecords = append(domainRecords, record)
	}

	return &GetDomainRecordsResponse{
		RequestId:     response.RequestId,
		Total:         response.TotalCount,
		DomainRecords: domainRecords,
	}, nil
}

func AddDomainRecord(r Record) (*DomainRecordResponse, error) {
	request := alidns.CreateAddDomainRecordRequest()
	request.Scheme = "https"
	request.DomainName = r.DomainName
	request.RR = r.RR
	request.Type = r.Type
	request.Value = r.Value

	response, err := Client.AddDomainRecord(request)
	if err != nil {
		return nil, err
	}

	return &DomainRecordResponse{
		RecordId:  response.RecordId,
		RequestId: response.RequestId,
	}, nil
}

func DeleteDomainRecord(r Record) (*DomainRecordResponse, error) {
	request := alidns.CreateDeleteDomainRecordRequest()
	request.Scheme = "https"
	request.RecordId = r.RecordId
	response, err := Client.DeleteDomainRecord(request)
	if err != nil {
		return nil, err
	}

	return &DomainRecordResponse{
		RecordId:  response.RecordId,
		RequestId: response.RequestId,
	}, nil
}

func UpdateDomainRecord(r Record) (*DomainRecordResponse, error) {
	request := alidns.CreateUpdateDomainRecordRequest()
	request.Scheme = "https"

	request.RecordId = r.RecordId
	request.RR = r.RR
	request.Type = r.Type
	request.Value = r.Value

	response, err := Client.UpdateDomainRecord(request)
	if err != nil {
		return nil, err
	}

	return &DomainRecordResponse{
		RecordId:  response.RecordId,
		RequestId: response.RequestId,
	}, nil
}

func SetDomainRecordStatus(r Record) (*DomainRecordResponse, error) {
	request := alidns.CreateSetDomainRecordStatusRequest()
	request.Scheme = "https"

	request.RecordId = r.RecordId
	request.Status = r.Status

	response, err := Client.SetDomainRecordStatus(request)
	if err != nil {
		return nil, err
	}

	return &DomainRecordResponse{
		RecordId:  response.RecordId,
		RequestId: response.RequestId,
	}, nil
}
