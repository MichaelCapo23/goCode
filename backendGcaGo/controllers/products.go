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

//GetProducts gets a list of products depending on the type given
func (c Controller) GetProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Product

		//define the params, required params and headers
		paramsReq, params := []string{"type"}, []string{}
		headers := [2]string{"token", "store_id"}

		//get the token/store_id, make the header map and add the token/store_id to the header Map
		token := r.Header.Get("token")
		storeID := r.Header.Get("store_id")
		headerMap := make(map[string]string)

		//add the token and store_id values to the header map
		headerMap["token"] = token
		headerMap["store_id"] = storeID

		//parse the form
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

		//make the slice response model
		res = []models.Product{}

		repo := repository.ProductRepo{}

		//run the GetProductsRepo method on the repo struct
		res, err := repo.GetProductsRepo(db, product, res, headerMap, paramsMap)
		if err != nil {
			err := models.Error{"Unable to get products"}
			utils.SendError(w, 500, err)
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
			utils.SendError(w, 500, err)
			return
		}

		//make the required/optional paramaters, and the headers
		paramsReq, params, headers := []string{}, []string{}, [2]string{"token", "store_id"}

		//get the token/store_id, make the header map and add the token/store_id to the header Map
		token := r.Header.Get("token")
		storeID := r.Header.Get("store_id")
		headerMap := make(map[string]string)

		//add the store_id and token to the headers map
		headerMap["store_id"] = storeID
		headerMap["token"] = token

		//parse the form
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

		//run the GetProductRepo on the repo struct
		res, err := repo.GetProductRepo(db, product, id)
		if err != nil {
			err := models.Error{"Unable to get product"}
			utils.SendError(w, 500, err)
		}

		//encode our res slice and send it back to the front end
		json.NewEncoder(w).Encode(res)
	}
}

func (c Controller) AddProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//make the params, required params and headers and headers map
		paramsReq := []string{"product_name", "category_id", "manufacturer", "description", "priceArr"}
		params := []string{"image_url", "status", "recommended", "subcategory_id"}
		headers := [2]string{"token", "store_id"}
		headerMap := make(map[string]string)

		//get the values from the headers
		token := r.Header.Get("token")
		storeID := r.Header.Get("store_id")

		//store the header values into the headers map
		headerMap["store_id"] = storeID
		headerMap["token"] = token

		//parse the form
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

		//run the AddProductRepo on the repo struct
		res, err := repo.AddProductRepo(db, headerMap, paramsMap)
		if err != nil {
			err := models.Error{res}
			utils.SendError(w, 500, err)
			return
		}

		//encode our res slice and send it back to the front end
		utils.SendSuccess(w, res)
	}
}

func (c Controller) EditProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//define the required params, params, headers and the headers map
		paramsReq := []string{"id"}
		params := []string{"product_name", "category_id", "manufacturer", "description", "image_url", "status", "recommended", "subcategory_id"}
		headers := [2]string{"token", "store_id"}
		headerMap := make(map[string]string)

		//get the values from the headers store them
		token := r.Header.Get("token")
		storeID := r.Header.Get("store_id")

		//add the headers variables to the headers map
		headerMap["store_id"] = storeID
		headerMap["token"] = token

		//parse the form
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

		//run the EditProductRepo method on the repo struct
		res, err := repo.EditProductRepo(db, headerMap, paramsMap)
		if err != nil {
			errMsg := models.Error{res}
			utils.SendError(w, 400, errMsg)
			return
		}

		//encode our res slice and send it back to the front end
		utils.SendSuccess(w, res)
	}
}
