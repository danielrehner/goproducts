package mocks

import (
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type MockedScanOutput struct {
	dynamodbiface.DynamoDBAPI
	Resp    dynamodb.ScanOutput
	GetResp dynamodb.GetItemOutput
}

func (m MockedScanOutput) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return &m.GetResp, nil
}

func (m MockedScanOutput) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return &m.Resp, nil
}

type MockedSearchOutput struct {
	cloudsearchdomainiface.CloudSearchDomainAPI
	Resp cloudsearchdomain.SearchOutput
}

func (m MockedSearchOutput) Search(input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error) {
	return &m.Resp, nil
}
