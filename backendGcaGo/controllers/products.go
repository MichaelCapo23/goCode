package controllers

import (
	"backendGcaGo/driver"
	"backendGcaGo/models"
	"backendGcaGo/repository"
	"backendGcaGo/utils"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct{}

var res []models.Product

func (c Controller) GetProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Product

		params := []string{"type"}
		token := r.Header.Get("token")
		storeID := r.Header.Get("store_id")

		headers := [2]string{"token", "store_id"}
		headerMap := make(map[string]string)
		headerMap["store_id"] = storeID
		headerMap["token"] = token

		r.ParseMultipartForm(0)

		//check required body params are sent, check token and expiration on token
		missing, ok, auth, paramsMap := utils.CheckTokenAndParams(headers, headerMap, params, r, db)
		if !auth {
			err := models.Error{"Invalid session"}
			utils.SendError(w, 401, err)
			return
		} else if !ok {
			err := models.Error{"Missing " + missing + " parameter"}
			utils.SendError(w, 422, err)
			return
		}

		res = []models.Product{}

		repo := repository.ProductRepo{}

		res, err := repo.GetProductsRepo(db, product, res, headerMap, paramsMap)
		driver.LogFatal(err)

		//encode our res slice and send it back to the front end
		utils.SendSuccess(w, res)
	}
}

func (c Controller) GetProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Product

		//get the URL parameter id, check errors
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			err := models.Error{"Invalid query parameter"}
			utils.SendError(w, 400, err)
			return
		}

		params := []string{}
		token := r.Header.Get("token")
		storeID := r.Header.Get("store_id")

		headers := [2]string{"token", "store_id"}
		headerMap := make(map[string]string)
		headerMap["store_id"] = storeID
		headerMap["token"] = token

		r.ParseMultipartForm(0)

		//check required body params are sent, check token and expiration on token
		missing, ok, auth, _ := utils.CheckTokenAndParams(headers, headerMap, params, r, db)
		if !auth {
			err := models.Error{"Invalid session"}
			utils.SendError(w, 401, err)
			return
		} else if !ok {
			err := models.Error{"Missing " + missing + " parameter"}
			utils.SendError(w, 422, err)
			return
		}

		res = []models.Product{}
		repo := repository.ProductRepo{}

		res, err := repo.GetProductRepo(db, product, res, id)
		driver.LogFatal(err)

		//encode our res slice and send it back to the front end
		json.NewEncoder(w).Encode(res)
	}
}
