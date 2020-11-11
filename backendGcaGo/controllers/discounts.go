package controllers

import (
	"backendGcaGo/models"
	repository "backendGcaGo/repository/discounts"
	"backendGcaGo/utils"
	"database/sql"
	"net/http"
)

func (c Controller) GetDiscounts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//define the params, required params and headers
		paramsReq, params, headers := []string{}, []string{}, [2]string{"token", "store_id"}

		//get the token, make the header map and add the token to the header Map
		token, storeID, headerMap := r.Header.Get("token"), r.Header.Get("store_id"), make(map[string]string)
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

		res := []models.Discount{}

		repo := repository.DiscountRepo{}

		res, err := repo.GetProductsRepo(db, product, res, headerMap, paramsMap)
		if err != nil {
			err := models.Error{"Unable to get products"}
			utils.SendError(w, 400, err)
		}

		//encode our res slice and send it back to the front end
		utils.SendSuccess(w, res)
	}
}
