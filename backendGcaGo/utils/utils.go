package utils

import (
	"backendGcaGo/driver"
	"backendGcaGo/models"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

func SendError(w http.ResponseWriter, status int, err interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func CheckTokenAndParams(headers [2]string, h map[string]string, paramsReq []string, params []string, r *http.Request, db *sql.DB) (string, bool, bool, map[string]string) {
	//make params map
	paramsMap := make(map[string]string)

	// check the headers are set
	for _, c := range headers {
		ok := h[c]
		if ok == "" {
			return "", false, false, paramsMap
		}
	}

	//check the required params are set
	for _, p := range paramsReq {
		ok := r.FormValue(p)
		if ok == "" {
			return p, false, true, paramsMap
		}
		paramsMap[p] = ok
	}

	//check the params are set
	for _, p := range params {
		ok := r.FormValue(p)
		if ok != "" {
			paramsMap[p] = ok
		}
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

//AddProductPrice adds new rows to the prices table by the new product id
func AddProductPrice(tx *sql.Tx, c map[string]string, insertID int64, res string, ch chan bool) {
	cols, vals, valuesArr, stmt, ctx := "", "", make([]interface{}, 0), "", context.Background()

	//loop over each map in the priceArr, build parts of query
	for key, val := range c {
		cols += key + ","
		vals += "?,"
		valuesArr = append(valuesArr, val)
	}

	//add last insert id to query and values array here
	cols += "product_id,"
	vals += "?,"
	valuesArr = append(valuesArr, insertID)

	//remove the trailing "," from both variables
	cols = cols[:len(cols)-1]
	vals = vals[:len(vals)-1]

	//execute query, rollback if err found
	stmt = "INSERT INTO `prices` (" + cols + ") VALUES (" + vals + ")"
	_, err := tx.ExecContext(ctx, stmt, valuesArr...)
	if err != nil {
		tx.Rollback()
		ch <- false
	}
	ch <- true
}

//GetProductNameImg gets the `product_name`, `image_url` to add to res slice of structs
func GetProductNameImg(db *sql.DB, res []models.Favorite, m models.Favorite, i int, ch chan bool) {
	//create valuesArr to pass to query
	valuesArr := make([]interface{}, 0)
	valuesArr = append(valuesArr, m.ProductID)

	//make query
	stmt := "SELECT `product_name`, `image_url` FROM `products_master_list` WHERE `id` = ? "

	//create a nullString incase image url is null, then grab and return the string
	var imgURL sql.NullString

	err := db.QueryRow(stmt, valuesArr...).Scan(&m.ProductName, &imgURL)
	m.ImageURL = imgURL.String
	res[i] = m

	if err != nil {
		ch <- false
	}
	ch <- true
}
