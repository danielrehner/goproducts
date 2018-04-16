package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/goproducts/database"
	"github.com/goproducts/dto"
)

func setupData(svc *database.DB) {
	result, _ := svc.Client.ListTables(&dynamodb.ListTablesInput{})
	tableNames := result.TableNames
	for !contains(tableNames, "GoApp.Products") {
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
		TableName: aws.String("GoApp.Products"),
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

	var err error
	_, err = svc.Client.CreateTable(createTableInput)

	if err != nil {
		fmt.Println("Error on CreateTable:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Created Table GoApp.Products")
}

func populateTable(svc *database.DB) {
	products := []dto.Product{
		dto.Product{ID: "1", Price: 3.74, Title: "Chocolate Chip Cookies"},
		dto.Product{ID: "2", Price: 8.21, Title: "Peanut Butter Cookies"},
	}
	for _, product := range products {
		av, err := dynamodbattribute.MarshalMap(product)

		if err != nil {
			fmt.Println("MarshalMap Error:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String("GoApp.Products"),
		}

		_, err = svc.Client.PutItem(input)

		if err != nil {
			fmt.Println("Got error calling PutItem:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Put product in GoApp.Products: " + product.Title)
	}
}

func verifyData(svc *database.DB) {
	getItemResult, err := svc.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("GoApp.Products"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String("1"),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	retrievedProduct := dto.Product{}

	err = dynamodbattribute.UnmarshalMap(getItemResult.Item, &retrievedProduct)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

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
