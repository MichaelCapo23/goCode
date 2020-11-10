package repository

import (
	"backendGcaGo/driver"
	"backendGcaGo/models"
	"context"
	"database/sql"
	"encoding/json"
)

//AddProductRepo function to add new products to the master_products_list table
func (ps ProductRepo) AddProductRepo(db *sql.DB, product models.Product, h map[string]string, p map[string]string) (string, error) {
	//make variables to build query and give values to prepared statement
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		driver.LogFatal(err)
	}
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

	stmt := "INSERT INTO products_master_list (" + cols + ") VALUES (" + vals + ")"

	rows, err2 := tx.ExecContext(ctx, stmt, valuesArr...)
	if err2 != nil {
		tx.Rollback()
		return res, err2
	}

	insertID, err3 := rows.LastInsertId()
	if err != nil {
		tx.Rollback()
		return res, err3
	}

	//add the prices to the prices table section
	//Unmarshal the JSON array of maps set new variable with the results
	priceArr := make([]map[string]string, 0)
	json.Unmarshal([]byte(p["priceArr"]), &priceArr)
	for _, c := range priceArr {
		//define the cols and values to build teh query
		cols, vals, valuesArr, stmt = "", "", nil, ""

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
			return res, err
		}
	}

	// Commit the change if all queries ran successfully
	err = tx.Commit()
	if err != nil {
		driver.LogFatal(err)
	}

	//change the res text now that no errors occurred
	res = "Successfully added a new product"
	return res, nil
}
