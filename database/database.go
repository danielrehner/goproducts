package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/goproducts/config"
	"github.com/goproducts/errors"
)

var svc *dynamodb.DynamoDB
var searchClient *cloudsearchdomain.CloudSearchDomain
var sess *session.Session

// DB struct enables data access and allows mocking for tests.
type DB struct {
	Client dynamodbiface.DynamoDBAPI
}

// Search struct enables search access and allows mocking for tests.
type Search struct {
	Client cloudsearchdomainiface.CloudSearchDomainAPI
}

// GetClient provides a dynamodb client.
func GetClient() dynamodbiface.DynamoDBAPI {
	if svc != nil {
		return svc
	}

	svc = dynamodb.New(GetSession(), &aws.Config{
		Endpoint: aws.String(config.GetString("dynamodb.endpoint")),
	})

	return svc
}

// GetSearchClient provides a search client for CloudSearch.
func GetSearchClient() cloudsearchdomainiface.CloudSearchDomainAPI {
	if searchClient != nil {
		return searchClient
	}

	searchClient = cloudsearchdomain.New(GetSession(), &aws.Config{
		Endpoint: aws.String(config.GetString("cloudsearch.endpoint")),
	})

	return searchClient
}

// GetSession provides an AWS session.
func GetSession() *session.Session {
	if sess != nil {
		return sess
	}

	creds := credentials.NewSharedCredentials(config.GetString("aws.credentials_file_path"), config.GetString("aws.credentials_name"))
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.GetString("aws.region")),
		Credentials: creds,
	})
	errors.HandleIfError(err)

	return sess
}
