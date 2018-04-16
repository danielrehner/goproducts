package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goproducts/database"
	"github.com/goproducts/server"
)

var r *gin.Engine
var db *database.DB

func main() {
	db := &database.DB{
		Client: database.GetClient(),
	}
	search := &database.Search{
		Client: database.GetSearchClient(),
	}
	setupData(db)
	r := server.SetupRouter(db, search)
	r.Run()

	// - Use go aws sdk or cli or cloudformation template file? both?
	// Scrape / find online sample data?
	// [x] Sort by price, descending by default
	// Hook up CloudSearch
	// DAX?
	// Add tests
	// Add benchmark tests?
	// Add example tests?

}
