package controllers

import (
	"backendGcaGo/driver"
	"backendGcaGo/models"
	"backendGcaGo/repository"
	"database/sql"
	"encoding/json"
	"net/http"
)

type Controller struct{}

var res []models.Product

func (c Controller) GetProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Product

		res = []models.Product{}

		repo := repository.ProductRepo{}

		res, err := repo.GetProductsRepo(db, product, res)
		driver.LogFatal(err)

		//encode our res slice and send it back to the front end
		json.NewEncoder(w).Encode(res)
	}
}
