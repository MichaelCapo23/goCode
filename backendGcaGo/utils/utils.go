package utils

import (
	"backendGcaGo/driver"
	"backendGcaGo/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func SendError(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func CheckTokenAndParams(headers [2]string, h map[string]string, params []string, r *http.Request, db *sql.DB) (string, bool, bool, map[string]string) {
	//make params map
	paramsMap := make(map[string]string)

	// check the headers are set
	for _, c := range headers {
		ok := h[c]
		if ok == "" {
			fmt.Println("in")
			return "", false, false, paramsMap
		}
	}

	//check the params are set
	for _, p := range params {
		ok := r.FormValue(p)
		if ok == "" {
			return p, false, true, paramsMap
		}
		paramsMap[p] = ok
	}

	rows, err := db.Query("SELECT `account_id`, `expiration` FROM `session` WHERE `token` = ?", h["token"])
	if err != nil {
		return "", true, false, paramsMap
	}

	count, authModel := checkCount(rows)

	//check rows given from token query
	if count != 1 {
		return "", true, false, paramsMap
	}

	//check the expiration of query
	ex := authModel.Expiration
	t := time.Now()
	c := t.Format("2006-01-02 15:04:05")

	if c > ex {
		fmt.Println("time err")
		return "", true, false, paramsMap
	}

	return "", true, true, paramsMap
}

func checkCount(rows *sql.Rows) (int, models.AuthModel) {
	authModel := models.AuthModel{}
	count := 0
	for rows.Next() {
		err := rows.Scan(&authModel.AccountID, &authModel.Expiration)
		driver.LogFatal(err)
		count++
	}
	return count, authModel
}
