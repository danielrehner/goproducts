package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goproducts/dto"
	"github.com/goproducts/search"
	"github.com/swaggo/swag/example/celler/httputil"
)

// ShowProduct godoc
// @Summary Shows the details of a product
// @Description Uses DynamoDB to retrieve details for a given product
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Success 200 {object} dto.ProductResponse
// @Router /products/show/:id [get]
func (c *Controller) ShowProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := search.ShowProduct(&c.DB, id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, &dto.ProductResponse{Data: product})
}

// SearchProducts godoc
// @Summary Searches for products using CloudSearch
// @Description Uses CloudSearch to find products by Title
// @Accept  json
// @Produce  json
// @Param q query string true "Search Term"
// @Success 200 {object} dto.ProductSearchResult
// @Router /products/search?q={q} [get]
func (c *Controller) SearchProducts(ctx *gin.Context) {
	searchQuery := ctx.Query("q")
	ctx.JSON(http.StatusOK, &dto.ProductSearchResult{
		Data: search.ProductSearch(&c.Search, searchQuery),
	})
}

// ScanProducts godoc
// @Summary Scans for products using DynamoDB
// @Description Uses DynamoDB to find products by Title
// @Accept  json
// @Produce  json
// @Param q query string true "Search Term"
// @Success 200 {object} dto.ProductSearchResult
// @Router /products/scan?q={q} [get]
func (c *Controller) ScanProducts(ctx *gin.Context) {
	searchQuery := ctx.Query("q")
	ctx.JSON(http.StatusOK, &dto.ProductSearchResult{
		Data: search.ProductScan(&c.DB, searchQuery),
	})
}
