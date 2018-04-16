package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goproducts/database"
	"github.com/goproducts/search"
)

// SetupRouter initializes Gin router and routes
func SetupRouter(db *database.DB, dbsearch *database.Search) *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/products/scan", func(c *gin.Context) {
		searchQuery := c.Query("q")
		c.JSON(200, gin.H{
			"data": search.ProductScan(db, searchQuery),
		})
	})

	r.GET("/products/search", func(c *gin.Context) {
		searchQuery := c.Query("q")
		c.JSON(200, gin.H{
			"data": search.ProductSearch(dbsearch, searchQuery),
		})
	})

	return r
}
