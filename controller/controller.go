package controller

import (
	"github.com/goproducts/database"
)

// Controller is the main controller struct upon which other controller functions build.
type Controller struct {
	database.DB
	database.Search
}

// New creates a new controller struct and sets the appropriate data access objects.
func New(db *database.DB, dbsearch *database.Search) *Controller {
	return &Controller{
		DB:     *db,
		Search: *dbsearch,
	}
}
