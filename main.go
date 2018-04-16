package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/goproducts/config"
	"github.com/goproducts/database"
	"github.com/goproducts/server"
)

var r *gin.Engine
var db *database.DB
var search *database.Search

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
