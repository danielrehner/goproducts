package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/goproducts/database"
	"github.com/goproducts/dto"
	"github.com/goproducts/server"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	response := parseSimpleJSONResponse(t, doGetRequest(t, "/"))

	value, exists := response["hello"]
	assert.True(t, exists)

	body := gin.H{
		"hello": "world",
	}
	assert.Equal(t, body["hello"], value)
}

func TestPing(t *testing.T) {
	response := parseSimpleJSONResponse(t, doGetRequest(t, "/ping"))

	value, exists := response["message"]
	assert.True(t, exists)

	body := gin.H{
		"message": "pong",
	}
	assert.Equal(t, body["message"], value)
}

func TestShowProduct(t *testing.T) {
	assert.Nil(t, nil)
	response := doGetRequest(t, "/api/v1/products/show/1")
	assert.NotNil(t, response)
	var responseProduct dto.ProductResponse
	err := json.Unmarshal([]byte(response.Body.String()), &responseProduct)
	assert.Nil(t, err)
	product := responseProduct.Data
	assert.NotNil(t, product)
	assert.Equal(t, "1", product.ID)
	assert.Equal(t, "Cookies", product.Title)
	assert.Equal(t, 1.00, product.Price)
}

func TestProductScan(t *testing.T) {
	response := doProductScan(t, "Cookies")
	product := response[0]
	assert.Equal(t, "1", product.ID)
	assert.Equal(t, "Cookies", product.Title)
	assert.Equal(t, 1.00, product.Price)
}

func TestProductSearch(t *testing.T) {
	doProductSearch(t, "Cookies")
	response := doProductSearch(t, "Cookies")
	product := response[0]
	assert.Equal(t, "1", product.ID)
	assert.Equal(t, "Cookies", product.Title)
	assert.Equal(t, 1.00, product.Price)
}

func doRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func doGetRequest(t *testing.T, path string) *httptest.ResponseRecorder {
	db := &database.DB{
		Client: mockedScanOutput{
			Resp: dynamodb.ScanOutput{
				Count: aws.Int64(1),
				Items: []map[string]*dynamodb.AttributeValue{
					{
						"ID": {
							S: aws.String("1"),
						},
						"Title": {
							S: aws.String("Cookies"),
						},
						"Price": {
							N: aws.String("1.00"),
						},
					},
				},
			},
			GetResp: dynamodb.GetItemOutput{
				Item: map[string]*dynamodb.AttributeValue{
					"ID": {
						S: aws.String("1"),
					},
					"Title": {
						S: aws.String("Cookies"),
					},
					"Price": {
						N: aws.String("1.00"),
					},
				},
			},
		},
	}
	dbsearch := &database.Search{
		Client: mockedSearchOutput{
			Resp: cloudsearchdomain.SearchOutput{
				Hits: &cloudsearchdomain.Hits{
					Hit: []*cloudsearchdomain.Hit{
						&cloudsearchdomain.Hit{
							Id: aws.String("1"),
							Fields: map[string][]*string{
								"id":    {aws.String("1")},
								"title": {aws.String("Cookies")},
								"price": {aws.String("1.00")},
							},
						},
					},
				},
			},
		},
	}
	r := server.SetupRouter(db, dbsearch)
	w := doRequest(r, "GET", path)
	assert.Equal(t, http.StatusOK, w.Code)
	return w
}

func parseSimpleJSONResponse(t *testing.T, w *httptest.ResponseRecorder) map[string]string {
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	return response
}

func doProductScan(t *testing.T, searchTerm string) []dto.Product {
	w := doGetRequest(t, "/api/v1/products/scan?q="+searchTerm)
	var response map[string][]dto.Product
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	value, exists := response["data"]
	assert.True(t, exists)
	return value
}

func doProductSearch(t *testing.T, searchTerm string) []dto.Product {
	w := doGetRequest(t, "/api/v1/products/search?q="+searchTerm)
	var response map[string][]dto.Product
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	value, exists := response["data"]
	assert.True(t, exists)
	return value
}
