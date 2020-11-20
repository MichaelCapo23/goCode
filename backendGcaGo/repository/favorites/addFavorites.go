package repository

import (
	"context"
	"database/sql"
	"errors"
)

//AddFavoritesRepo adds a new favorite row in the favorites table. takes app_suer_id, store_id and product_id
func (fs FavoritesRepo) AddFavoritesRepo(db *sql.DB, h map[string]string, p map[string]string) error {
	//creat new tx variable to hold a reference to *sql.Tx to start a new transaction in the db, only commit after error have been checked for.
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//create valuesArr to pass to the query
	valuesArr := make([]interface{}, 0)

	//append the values to the slice of interfaces to add to the perpared statement (next query checks for any rows that match this data if any found, return out of transaction)
	valuesArr = append(valuesArr, p["app_user_id"], p["product_id"], h["store_id"])

	//make new variable to store the count from the select query
	var rowsAffected string

	//build the query to check for a row that has matching data to see if it already exists.
	stmt := "SELECT COUNT(*) AS `count` FROM `favorites` WHERE `app_user_id` = ? AND `product_id` = ? AND `store_id` = ?"

	//run the query, scan the value (count) to the rowsAffected variable of type string
	err = db.QueryRow(stmt, valuesArr...).Scan(&rowsAffected)

	//check for errors/check for the rows affected
	if err != nil {
		return err
	} else if rowsAffected != "0" {
		res := errors.New("Matching favorite row found")
		return res
	}

	//build the query, valuesArr doesn't need to be updated, data inside should also work the the insert query (query adds a new row to the favorites table)
	stmt = "INSERT INTO `favorites` (`app_user_id`, `product_id`, `store_id`) VALUES (?,?,?)"

	//execute query don't commit the transaction yet
	_, err2 := tx.ExecContext(ctx, stmt, valuesArr...)

	//check for errors rollback transaction if error found
	if err2 != nil {
		tx.Rollback()
		return err2
	}

	//commit the transaction, check for error
	err = tx.Commit()
	if err != nil {
		return err
	}

	//return nil to show no errors were found
	return nil
}
