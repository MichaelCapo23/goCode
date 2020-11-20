package controllers

import (
	"backendGcaGo/models"
	repository "backendGcaGo/repository/favorites"
	"backendGcaGo/utils"
	"database/sql"
	"net/http"
)

//GetFavorites gets all of the favorites for the app user if given
func (c Controller) GetFavorites(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var favorite models.Favorite

		//define the params, required params and headers
		paramsReq, params := []string{"app_user_id"}, []string{}
		headers := [2]string{"token", "store_id"}

		//get the token, make the header map and add the token to the header Map
		token := r.Header.Get("token")
		storeID := r.Header.Get("store_id")
		headerMap := make(map[string]string)
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

		//make the response variable (slice of type models.Favorite)
		res := []models.Favorite{}

		repo := repository.FavoritesRepo{}

		//run the GetFavoritesRepo method on the repo struct
		res, err := repo.GetFavoritesRepo(db, favorite, res, headerMap, paramsMap)
		if err != nil {
			// err := models.Error{"Error getting Favorites"}
			errorMSG := err.Error()
			utils.SendError(w, 500, errorMSG)
			return
		}
		utils.SendSuccess(w, res)
	}
}

func (c Controller) AddFavorites(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//define the params, required params and headers
		paramsReq, params := []string{"app_user_id", "product_id"}, []string{}
		headers := [2]string{"token", "store_id"}

		//get the token, make the header map and add the token to the header Map
		token := r.Header.Get("token")
		storeID := r.Header.Get("store_id")
		headerMap := make(map[string]string)
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

		repo := repository.FavoritesRepo{}
		err := repo.AddFavoritesRepo(db, headerMap, paramsMap)
		if err != nil {
			err := models.Error{"Error adding new favorite"}
			utils.SendError(w, 500, err)
			return
		}
		utils.SendSuccess(w, "Successfully added new favorite")
	}
}

func (c Controller) EditFavorites(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paramsReq, params := []string{"app_user_id", "id", "product_id"}, []string{}
		headers := [2]string{"token", "store_id"}

		//get the token, make the header map and add the token to the header Map
		token := r.Header.Get("token")
		storeID := r.Header.Get("store_id")
		headerMap := make(map[string]string)
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

		repo := repository.FavoritesRepo{}
		err := repo.EditFavoritesRepo(db, headerMap, paramsMap)
		if err != nil {
			// err := models.Error{"Error updating new favorite"}
			utils.SendError(w, 400, err.Error())
			return
		}
		utils.SendSuccess(w, "Successfully edited favorite")
	}
}
