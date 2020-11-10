package controllers

import (
	"backendGcaGo/models"
	repository "backendGcaGo/repository/products"
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

		//define the params, required params and headers
		paramsReq, params := []string{"type"}, []string{}
		headers := [2]string{"token"}

		//get the token, make the header map and add the token to the header Map
		token := r.Header.Get("token")
		headerMap := make(map[string]string)
		headerMap["token"] = token

		r.ParseMultipartForm(0)

		//check required body params are sent, check token and expiration on token
		missing, ok, auth, paramsMap := utils.CheckTokenAndParams(headers, headerMap, paramsReq, params, r, db)
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
		if err != nil {
			err := models.Error{"Unable to get products"}
			utils.SendError(w, 400, err)
		}

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

		paramsReq, params := []string{}, []string{}
		token := r.Header.Get("token")
		storeID := r.Header.Get("store_id")

		headers := [2]string{"token", "store_id"}
		headerMap := make(map[string]string)
		headerMap["store_id"] = storeID
		headerMap["token"] = token

		r.ParseMultipartForm(0)

		//check required body params are sent, check token and expiration on token
		missing, ok, auth, _ := utils.CheckTokenAndParams(headers, headerMap, paramsReq, params, r, db)
		if !auth {
			err := models.Error{"Invalid session"}
			utils.SendError(w, 401, err)
			return
		} else if !ok {
			err := models.Error{"Missing " + missing + " parameter"}
			utils.SendError(w, 422, err)
			return
		}

		repo := repository.ProductRepo{}

		res, err := repo.GetProductRepo(db, product, id)
		if err != nil {
			err := models.Error{"Unable to get product"}
			utils.SendError(w, 400, err)
		}

		//encode our res slice and send it back to the front end
		json.NewEncoder(w).Encode(res)
	}
}

func (c Controller) AddProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Product

		paramsReq := []string{"product_name", "category_id", "manufacturer", "description", "priceArr"}
		params := []string{"image_url", "status", "recommended", "subcategory_id"}
		token := r.Header.Get("token")
		storeID := r.Header.Get("store_id")

		headers := [2]string{"token", "store_id"}
		headerMap := make(map[string]string)
		headerMap["store_id"] = storeID
		headerMap["token"] = token

		r.ParseMultipartForm(0)

		//check required body params are sent, check token and expiration on token
		missing, ok, auth, paramsMap := utils.CheckTokenAndParams(headers, headerMap, paramsReq, params, r, db)
		if !auth {
			err := models.Error{"Invalid session"}
			utils.SendError(w, 401, err)
			return
		} else if !ok {
			err := models.Error{"Missing " + missing + " parameter"}
			utils.SendError(w, 422, err)
			return
		}

		repo := repository.ProductRepo{}

		//make default response, will be updated later
		res, err := repo.AddProductRepo(db, product, headerMap, paramsMap)
		if err != nil {
			err := models.Error{res}
			utils.SendError(w, 400, err)
			return
		}

		//encode our res slice and send it back to the front end
		utils.SendSuccess(w, res)
	}
}
