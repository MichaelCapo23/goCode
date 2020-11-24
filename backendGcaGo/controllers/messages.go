package controllers

import (
	"backendGcaGo/models"
	repository "backendGcaGo/repository/messages"
	"backendGcaGo/utils"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func (c Controller) GetMessages(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//define the params, required params and headers
		message, paramsReq, params, headers := models.Message{}, []string{"type"}, []string{}, [2]string{"token", "store_id"}

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

		res := []models.Message{}

		repo := repository.MessagesRepo{}

		res, err := repo.GetMessagesRepo(db, message, res, paramsMap, headerMap)
		if err != nil {
			// msg := models.Error{"Unable to get messages"}
			msg := err.Error()
			utils.SendError(w, 500, msg)
			return
		}
		utils.SendSuccess(w, res)
	}
}

//GetMessage gets a single message based off the id in the route
func (c Controller) GetMessage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//define the params, required params and headers
		message, paramsReq, params, headers := models.Message{}, []string{}, []string{}, [2]string{"token", "store_id"}

		//store the vars into variable
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			err := models.Error{"Invalid query parameter"}
			utils.SendError(w, 400, err)
			return
		}

		//get the token, make the header map and add the token to the header map
		token, storeID, headerMap := r.Header.Get("token"), r.Header.Get("store_id"), make(map[string]string)

		//store the values in the map
		headerMap["token"] = token
		headerMap["store_id"] = storeID

		//parse the form
		r.ParseMultipartForm(0)

		//check required body params are sent, check token and expiration on token
		missing, ok, auth, _ := utils.CheckTokenAndParams(headers, headerMap, paramsReq, params, r, db)
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

		repo := repository.MessagesRepo{}

		res, err := repo.GetMessageRepo(db, message, headerMap, id)

		if err != nil {
			msg := models.Error{"Unable to get message"}
			// msg := err.Error()
			utils.SendError(w, 500, msg)
			return
		}

		//send back the message
		utils.SendSuccess(w, res)
	}
}

func (c Controller) AddMessage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//define the params, required params and headers
		paramsReq, params, headers := []string{"type", "message", "app_user_id"}, []string{}, [2]string{"token", "store_id"}

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

		repo := repository.MessagesRepo{}

		err := repo.AddMessageRepo(db, paramsMap, headerMap)

		if err != nil {
			msg := models.Error{"Unable to add new message"}
			// msg := err.Error()
			utils.SendError(w, 500, msg)
			return
		}

		utils.SendSuccess(w, "Successfully added new message")
	}
}

//EditMessage edits message row based off id passed in the parameters
func (c Controller) EditMessage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//define the params, required params and headers
		paramsReq, params, headers := []string{"id"}, []string{}, [2]string{"token", "store_id"}

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

		repo := repository.MessagesRepo{}

		err := repo.EditMessageRepo(db, paramsMap, headerMap)

		if err != nil {
			// msg := models.Error{"Unable to edit message"}
			msg := err.Error()
			utils.SendError(w, 500, msg)
			return
		}

		utils.SendSuccess(w, "Successfully updated message")

	}
}
