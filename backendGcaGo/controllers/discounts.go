package controllers

import (
	"backendGcaGo/models"
	repository "backendGcaGo/repository/discounts"
	"backendGcaGo/utils"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func (c Controller) GetDiscounts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//define the params, required params and headers
		discount, paramsReq, params, headers := models.Discount{}, []string{}, []string{}, [2]string{"token", "store_id"}

		//get the token, make the header map and add the token to the header Map
		token, storeID, headerMap := r.Header.Get("token"), r.Header.Get("store_id"), make(map[string]string)
		headerMap["token"] = token
		headerMap["store_id"] = storeID

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

		res := [][]models.Discount{}

		repo := repository.DiscountRepo{}

		res, err := repo.GetDiscountsRepo(db, discount, res, headerMap, paramsMap)
		if err != nil {
			err := models.Error{"Unable to get products"}
			utils.SendError(w, 400, err)
			return
		}

		//encode our res slice and send it back to the front end
		utils.SendSuccess(w, res)
	}
}

func (c Controller) GetDiscount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//define the params, required params and headers
		discount, paramsReq, params, headers, vars := models.Discount{}, []string{}, []string{}, [2]string{"token", "store_id"}, mux.Vars(r)
		id, ok := vars["id"]

		//get the token, make the header map and add the token to the header Map
		token, storeID, headerMap := r.Header.Get("token"), r.Header.Get("store_id"), make(map[string]string)
		headerMap["token"] = token
		headerMap["store_id"] = storeID

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

		repo := repository.DiscountRepo{}

		res, err := repo.GetDiscountRepo(db, discount, headerMap, id)
		if err != nil {
			err := models.Error{"Unable to get products"}
			utils.SendError(w, 400, err)
			return
		}

		//encode our res slice and send it back to the front end
		utils.SendSuccess(w, res)
	}
}

func (c Controller) AddDiscount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//define the params, required params and headers
		paramsReq, params, headers := []string{"start_date", "end_date", "product_id", "discounts"}, []string{}, [2]string{"token", "store_id"}

		//get the token, make the header map and add the token to the header Map
		token, storeID, headerMap := r.Header.Get("token"), r.Header.Get("store_id"), make(map[string]string)
		headerMap["token"] = token
		headerMap["store_id"] = storeID

		//parse the request body
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

		res := "Successfully updated discounts"

		repo := repository.DiscountRepo{}

		msg, err := repo.AddDiscountRepo(db, headerMap, paramsMap)
		if err != nil {
			err := models.Error{msg}
			utils.SendError(w, 400, err)
			return
		}

		utils.SendSuccess(w, res)

	}
}
