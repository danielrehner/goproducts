package data

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/goproducts/config"
	"github.com/goproducts/database"
	"github.com/goproducts/dto"
	"github.com/goproducts/errors"
)

// setupData creates the products table, populates it with an initial data set, and makes sure a product is present.
func SetupData(svc *database.DB) {
	result, _ := svc.Client.ListTables(&dynamodb.ListTablesInput{})
	tableNames := result.TableNames

	waitCounter := 0
	for !contains(tableNames, config.GetString("dynamodb.productsTableName")) {
		createTable(svc)
		result, _ := svc.Client.ListTables(&dynamodb.ListTablesInput{})
		tableNames = result.TableNames
		time.Sleep(2 * time.Second)
		waitCounter++
		if waitCounter >= 10 {
			break
		}
	}

	populateTable(svc)
	verifyData(svc)
}

// createTable creates the configured products table.
func createTable(svc *database.DB) {

	createTableInput := &dynamodb.CreateTableInput{
		TableName: aws.String(config.GetString("dynamodb.productsTableName")),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
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

// populateTable puts a small set of initial documents in the products table.
func populateTable(svc *database.DB) {

	csvFile, _ := os.Open(config.GetString("data.productsFile"))
	reader := csv.NewReader(bufio.NewReader(csvFile))
	products := []dto.Product{}
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		price, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			price = 1.0
		}
		products = append(products, dto.Product{
			ID:    line[0],
			Title: line[1],
			Price: price,
		})
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

// verifyData makes sure the first product is present in the table.
func verifyData(svc *database.DB) {
	getItemResult, err := svc.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(config.GetString("dynamodb.productsTableName")),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
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
}

// contains returns true if the name is found in aRange and returns false otherwise.
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
