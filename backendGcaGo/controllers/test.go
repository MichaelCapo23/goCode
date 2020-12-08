package controllers

import (
	"backendGcaGo/models"
	repository "backendGcaGo/repository/test"
	"backendGcaGo/utils"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

//Test is a method to test new packages and try to add new tools to my arsenal
func (c Controller) Test(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//define the params, required params and headers
		paramsReq, params, headers := []string{}, []string{}, [2]string{"token", "store_id"}

		//get the params from url, extract the id to use in the next function call
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			err := models.Error{"Invalid query parameter"}
			utils.SendError(w, 500, err)
			return
		}

		//make a map to store the headers
		headerMap := make(map[string]string)

		//get the token and store_id from the headers
		token := r.Header.Get("token")
		storeID := r.Header.Get("store_id")

		//make new key values pairs in the header map to store the variables
		headerMap["token"] = token
		headerMap["store_id"] = storeID

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

		repo := repository.TestRepo{}

		err := repo.TestMutexIsolationLevels(db, headerMap, id)
		if err != nil {
			utils.SendError(w, 500, err.Error())
		}
		utils.SendSuccess(w, "got info")
	}
}
