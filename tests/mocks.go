package tests

import (
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type mockedScanOutput struct {
	dynamodbiface.DynamoDBAPI
	Resp dynamodb.ScanOutput
}

func (m mockedScanOutput) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return &m.Resp, nil
}

type mockedSearchOutput struct {
	cloudsearchdomainiface.CloudSearchDomainAPI
	Resp cloudsearchdomain.SearchOutput
}

func (m mockedSearchOutput) Search(input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error) {
	return &m.Resp, nil
}
