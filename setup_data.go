package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/goproducts/config"
	"github.com/goproducts/database"
	"github.com/goproducts/dto"
	"github.com/goproducts/errors"
)

func setupData(svc *database.DB) {
	result, _ := svc.Client.ListTables(&dynamodb.ListTablesInput{})
	tableNames := result.TableNames
	for !contains(tableNames, config.GetString("dynamodb.productsTableName")) {
		createTable(svc)
		result, _ := svc.Client.ListTables(&dynamodb.ListTablesInput{})
		tableNames = result.TableNames
		time.Sleep(2 * time.Second)
	}

	populateTable(svc)
	verifyData(svc)
}

func createTable(svc *database.DB) {

	createTableInput := &dynamodb.CreateTableInput{
		TableName: aws.String(config.GetString("dynamodb.productsTableName")),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}

	_, err := svc.Client.CreateTable(createTableInput)
	errors.HandleIfError(err)

	fmt.Println("Created Table " + config.GetString("dynamodb.productsTableName"))
}

func populateTable(svc *database.DB) {
	products := []dto.Product{
		dto.Product{ID: "1", Price: 3.74, Title: "Chocolate Chip Cookies"},
		dto.Product{ID: "2", Price: 8.21, Title: "Peanut Butter Cookies"},
	}
	for _, product := range products {
		av, err := dynamodbattribute.MarshalMap(product)
		errors.HandleIfError(err)

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(config.GetString("dynamodb.productsTableName")),
		}

		_, err = svc.Client.PutItem(input)
		errors.HandleIfError(err)
		fmt.Println("Put product: " + product.Title)
	}
}

func verifyData(svc *database.DB) {
	getItemResult, err := svc.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(config.GetString("dynamodb.productsTableName")),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String("1"),
			},
		},
	})
	errors.HandleIfError(err)

	retrievedProduct := dto.Product{}

	err = dynamodbattribute.UnmarshalMap(getItemResult.Item, &retrievedProduct)
	errors.HandleIfError(err)

	if retrievedProduct.Title == "" {
		fmt.Println("Could not find Product 1")
		return
	}

	fmt.Println("Found product:")
	fmt.Println("Title:  ", retrievedProduct.Title)
	fmt.Println("ID: ", retrievedProduct.ID)
	fmt.Println("Price:  ", retrievedProduct.Price)
}

func contains(aRange []*string, name string) bool {
	valueFound := false

	for _, n := range aRange {
		if name == *n {
			valueFound = true
			break
		}
	}
	return valueFound
}
