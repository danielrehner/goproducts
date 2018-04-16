package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var svc *dynamodb.DynamoDB
var searchClient *cloudsearchdomain.CloudSearchDomain
var sess *session.Session

// DB struct to enable data access and mocking for tests
type DB struct {
	Client dynamodbiface.DynamoDBAPI
}

// Search struct to enable search access and mocking for tests
type Search struct {
	Client cloudsearchdomainiface.CloudSearchDomainAPI
}

// GetClient provides a dynamodb client
func GetClient() dynamodbiface.DynamoDBAPI {
	if svc == nil {
		svc = dynamodb.New(GetSession())
	}
	return svc
}

// GetSearchClient provides a search client for CloudSearch
func GetSearchClient() cloudsearchdomainiface.CloudSearchDomainAPI {
	if searchClient == nil {
		searchClient = cloudsearchdomain.New(GetSession(), &aws.Config{
			Endpoint: aws.String("search-goproducts-domain-kg3aqdlqa7sttwc2e5vobmhkki.us-west-2.cloudsearch.amazonaws.com"),
		})
	}
	return searchClient
}

// GetSession provides a session
func GetSession() *session.Session {
	if sess == nil {
		creds := credentials.NewSharedCredentials("/Users/rehnerd/.aws/credentials", "default")
		sess, _ = session.NewSession(&aws.Config{Region: aws.String("us-west-2"),
			Credentials: creds,
		},
		)
	}

	return sess
}
