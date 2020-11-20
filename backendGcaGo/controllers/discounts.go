package controllers

import (
	"backendGcaGo/models"
	repository "backendGcaGo/repository/discounts"
	"backendGcaGo/utils"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

//GetDiscounts gets an array of discounts based off the store_id given on the headers
func (c Controller) GetDiscounts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//define the params, required params and headers
		discount, paramsReq, params, headers := models.Discount{}, []string{}, []string{}, [2]string{"token", "store_id"}

		//get the token, make the header map and add the token to the header map
		token, storeID, headerMap := r.Header.Get("token"), r.Header.Get("store_id"), make(map[string]string)

		//store the values in the map
		headerMap["token"] = token
		headerMap["store_id"] = storeID

		//parse the form
		r.ParseMultipartForm(0)

		//check required body params are sent, check token and expiration on token
		missing, ok, auth, paramsMap := utils.CheckTokenAndParams(headers, headerMap, paramsReq, params, r, db)
		if !auth {
			//if auth is false the token is expired, send back invalid session response
			err := models.Error{"Invalid session"}
			utils.SendError(w, 401, err)
			return
		} else if !ok {
			err := models.Error{"Missing " + missing + " parameter"}
			utils.SendError(w, 422, err)
			return
		}

		//make the response. should be a slice of slices based off the parent discount. Groups together the sub discounts (different prices and quantities)
		res := [][]models.Discount{}

		//instantiate a new DiscountRepo struct to call the GetDiscountsRepo method on
		repo := repository.DiscountRepo{}

		//call the method GetDiscountsRepo pass the db, the discount data model, the response model, header and params maps
		res, err := repo.GetDiscountsRepo(db, discount, res, headerMap, paramsMap)
		if err != nil {
			errorMsg := err.Error()
			utils.SendError(w, 500, errorMsg)
			return
		}

		//encode our res slice and send it back to the front end
		utils.SendSuccess(w, res)
	}
}

//GetDiscount returns a single discount based off the id sent in the route
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

		//instantiate a new DiscountRepo struct to call the GetDiscountsRepo method on
		repo := repository.DiscountRepo{}

		res, err := repo.GetDiscountRepo(db, discount, headerMap, id)
		if err != nil {
			err := models.Error{"Unable to get products"}
			utils.SendError(w, 500, err)
			return
		}

		//encode our res slice and send it back to the front end
		utils.SendSuccess(w, res)
	}
}

// EditDiscount edits discounts based off the id in the discounts array of maps sent in the body parameters
func (c Controller) EditDiscount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//define the params, required params and headers
		paramsReq, params, headers := []string{"discounts"}, []string{}, [2]string{"token", "store_id"}

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

		//create the model for the response (just a simple successful message will do for now to show it was completed, will return custom error message if fails)
		res := "Successfully updated discounts"

		//instantiate a new DiscountRepo struct to call the GetDiscountsRepo method on
		repo := repository.DiscountRepo{}

		//run the EditDiscountRepo method on repo. check for errors, send back error and http status code
		msg, err := repo.EditDiscountRepo(db, headerMap, paramsMap)
		if err != nil {
			utils.SendError(w, 500, msg)
			return
		}

		//send back the response to
		utils.SendSuccess(w, res)
	}
}
