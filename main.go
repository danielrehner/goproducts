package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/goproducts/config"
	"github.com/goproducts/database"
	_ "github.com/goproducts/docs"
	"github.com/goproducts/server"
)

var r *gin.Engine

// @title GoProducts Example Project
// @version 1.0
// @description This is a sample for using Go, DynamoDB, and CloudSearch

// @contact.name Daniel Rehner
// @contact.url http://www.twitter.com/danielrehner
// @contact.email daniel.rehner@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
func main() {

	environment := flag.String("e", "development", "")
	flag.Parse()
	config.InitializeEnvironment(*environment)

	db := &database.DB{
		Client: database.GetClient(),
	}
	search := &database.Search{
		Client: database.GetSearchClient(),
	}

	setupData(db)
	r := server.SetupRouter(db, search)

	r.Run()
}
