package repository

import (
	"context"
	"database/sql"
	"errors"
)

func (fs FavoritesRepo) EditFavoritesRepo(db *sql.DB, h map[string]string, p map[string]string) error {
	//create tx variable to store the *sql.Tx for starting and committing transactions
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//make valuesArr to pass to the query
	valuesArr := make([]interface{}, 0)

	//fill the valuesArr with the values needed in the query (order is important)
	valuesArr = append(valuesArr, p["product_id"], p["app_user_id"], p["id"], h["store_id"])

	//build the query
	stmt := "UPDATE `favorites` SET `product_id` = ? WHERE (`app_user_id` = ? AND `id` = ?) AND `store_id` = ?"
	rows, err2 := tx.ExecContext(ctx, stmt, valuesArr...)

	//check for errors with the query, rollback query if errors found
	if err2 != nil {
		tx.Rollback()
		return err2
	}

	//get the rows affected to make sure the service ran and updated a row
	rowsAffected, err3 := rows.RowsAffected()
	if err3 != nil {
		tx.Rollback()
		return err3
	} else if rowsAffected != 1 {
		tx.Rollback()
		return errors.New("No rows updated")
	}

	//commit the transaction to the database
	err = tx.Commit()
	if err != nil {
		return err
	}

	//return nil to show no errors were found
	return nil
}
