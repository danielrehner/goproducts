package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goproducts/controller"
	"github.com/goproducts/database"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// SetupRouter initializes the gin router and routes.
func SetupRouter(db *database.DB, dbsearch *database.Search) *gin.Engine {
	r := gin.Default()

	c := controller.New(db, dbsearch)

	v1 := r.Group("/api/v1/products")
	{
		v1.GET("show/:id", c.ShowProduct)
		v1.GET("search", c.SearchProducts)
		v1.GET("scan", c.ScanProducts)
	}

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
