package repository

import (
	"backendGcaGo/driver"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (ps ProductRepo) EditProductRepo(db *sql.DB, h map[string]string, p map[string]string) (string, error) {
	res, cols, valuesArr := "Unable to edit product", "", make([]interface{}, 0)

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		driver.LogFatal(err)
	}

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

	stmt := "UPDATE `products_master_list` SET " + cols + " WHERE `store_id` = ? AND `id` = ?"

	row, err2 := tx.ExecContext(ctx, stmt, valuesArr...)
	if err2 != nil {
		tx.Rollback()
		return res, err2
	}

	affectedRows, err := row.RowsAffected()
	if err != nil || affectedRows != 1 {
		//rollback transaction create custom error message to show to the user
		tx.Rollback()
		myString := fmt.Sprintf("%v", affectedRows)
		res = "Number of rows affected:" + myString
		err = errors.New("Error updating product")
		return res, err
	}

	err = tx.Commit()
	if err != nil {
		driver.LogFatal(err)
	}

	res = "Successfully updated a new product"
	return res, nil
}
