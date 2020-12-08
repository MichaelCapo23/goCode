package repository

import (
	"context"
	"database/sql"
)

// AddMessageRepo adds a new message to the messages table. return nil if no errors found
func (ms MessagesRepo) AddMessageRepo(db *sql.DB, p map[string]string, h map[string]string) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//make the valuesArr to pass to the query
	valuesArr := make([]interface{}, 0)

	//append thr values to the slice to pass to the query
	valuesArr = append(valuesArr, h["store_id"], p["type"], p["app_user_id"], p["message"])

	//build the query
	stmt := "INSERT INTO `messages` (`store_id`, `type`, `app_user_id`, `message`) VALUES (?,?,?,?)"
	_, err = tx.ExecContext(ctx, stmt, valuesArr...)

	if err != nil {
		tx.Rollback()
		return err
	}

	//commit the transaction
	tx.Commit()
	return nil
}
