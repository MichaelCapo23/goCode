package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// EditProductRepo edits a row depending on the id sent
func (ps ProductRepo) EditProductRepo(db *sql.DB, h map[string]string, p map[string]string) (string, error) {
	//make a tx varibale of type *sql.Tx  to start a new transaction
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return "error making transaction", err
	}

	//make variables to build query and give values to prepared statement
	res, cols, valuesArr := "Unable to edit product", "", make([]interface{}, 0)

	//build part of the query by looping over the params sent. add it to the valuesArr to pass to the query
	for key, value := range p {
		if key != "id" {
			cols += key + " = ?,"
			valuesArr = append(valuesArr, value)
		}
	}

	//remove the trailing ","
	cols = cols[:len(cols)-1]

	//add store_id and the id  from the headers/params maps to the valuesArr
	valuesArr = append(valuesArr, h["store_id"])
	valuesArr = append(valuesArr, p["id"])

	//build the query
	stmt := "UPDATE `products_master_list` SET " + cols + " WHERE `store_id` = ? AND `id` = ?"

	//run the query and start the transaction. Check for errors if found rollback transaction
	row, err2 := tx.ExecContext(ctx, stmt, valuesArr...)
	if err2 != nil {
		tx.Rollback()
		return res, err2
	}

	//check the affected rows
	affectedRows, err := row.RowsAffected()

	//check for errors or if the rows affected isn't 1
	if err != nil || affectedRows != 1 {
		//rollback transaction create custom error message to show to the user
		tx.Rollback()
		myString := fmt.Sprintf("%v", affectedRows)
		res = "Number of rows affected:" + myString
		err = errors.New("Error updating product")
		return res, err
	}

	//commit the transaction
	err = tx.Commit()
	if err != nil {
		return "error trying to commit transaction", err
	}

	//send back a string showing it succeeded
	res = "Successfully updated a new product"
	return res, nil
}
