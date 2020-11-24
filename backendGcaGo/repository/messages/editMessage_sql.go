package repository

import (
	"context"
	"database/sql"
)

// EditMessageRepo edits message based off the id given
func (ms MessagesRepo) EditMessageRepo(db *sql.DB, p map[string]string, h map[string]string) error {
	//create tx variable to start a new transaction
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//make valuesArr to pass to the transaction
	valuesArr := make([]interface{}, 0)
	valuesArr = append(valuesArr, p["id"], h["store_id"])

	//build the query
	stmt := "UPDATE `messages` SET `resolved` = 1 WHERE `id` = ? AND `store_id` = ?"

	//execute the query
	_, err = tx.ExecContext(ctx, stmt, valuesArr...)

	//check for errors rollback query if found
	if err != nil {
		tx.Rollback()
		return err
	}

	//commit the transaction
	tx.Commit()
	return nil
}
