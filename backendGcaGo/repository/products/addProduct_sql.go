package repository

import (
	"backendGcaGo/utils"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
)

//AddProductRepo function to add new products to the master_products_list table
func (ps ProductRepo) AddProductRepo(db *sql.DB, h map[string]string, p map[string]string) (string, error) {
	//make a tx varibale of type *sql.Tx  to start a new transaction
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return "error making new transaction", err
	}

	//make variables to build query and give values to prepared statement
	cols, vals, valuesArr, res := "", "", make([]interface{}, 0), "Unable to add product"

	//add store_id to the params from the headers object
	p["store_id"] = h["store_id"]

	//loop over the params map, build the strings to add to the query string and add values to the valuesArr
	for key, val := range p {
		if key != "priceArr" {
			cols += key + ","
			vals += "?,"
			valuesArr = append(valuesArr, val)
		}
	}

	//remove the trailing "," from both variables
	cols = cols[:len(cols)-1]
	vals = vals[:len(vals)-1]

	//build the query with the cols and vals variables
	stmt := "INSERT INTO products_master_list (" + cols + ") VALUES (" + vals + ")"

	//run the query, dont commit yet. Check for errors, if any found rollback transaction
	rows, err2 := tx.ExecContext(ctx, stmt, valuesArr...)
	if err2 != nil {
		tx.Rollback()
		return res, err2
	}

	//get the last inserted id
	insertID, err3 := rows.LastInsertId()
	if err != nil {
		tx.Rollback()
		return res, err3
	}

	//add the prices to the prices table section
	//Unmarshal the JSON array of maps set new variable with the results
	priceArr := make([]map[string]string, 0)
	json.Unmarshal([]byte(p["priceArr"]), &priceArr)

	//create new channel of type bool to pass to goroutines
	ch := make(chan bool, len(priceArr))

	//loop over the priceArr and call a new go route to add each new price, pass the channel to it, if false returned, stop the process, don't commit the previous queries, return the error
	for _, c := range priceArr {
		//create a goroutines pass channel to add prices for the newly created product row
		go utils.AddProductPrice(tx, c, insertID, res, ch)
		if flag := <-ch; !flag {
			err := errors.New("Error adding new prices")
			res = "Error adding new prices"
			return res, err
		}
	}

	// Commit the change if all queries ran successfully
	err = tx.Commit()
	if err != nil {
		return "error committing transaction", err
	}

	//change the res text now that no errors occurred
	res = "Successfully added a new product"
	return res, nil
}
