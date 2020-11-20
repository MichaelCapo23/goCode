package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
)

var id chan string

//EditDiscountRepo updates discounts based of the id in the "discount" slice of maps
func (ds DiscountRepo) EditDiscountRepo(db *sql.DB, h map[string]string, p map[string]string) (string, error) {
	//make a new discounts variable and parse the slice of maps set the value of the new discounts variable
	discounts := make([]map[string]string, 0)
	json.Unmarshal([]byte(p["discounts"]), &discounts)

	ch := make(chan bool)

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	//start to update the/wait till the channel is filled to create new discounts
	for _, d := range discounts {
		go updateDiscount(tx, d, ch)
	}

	//loop over the length of discounts (the capacity for my channel, when new values is passed check if its false, send back error if false)
	for index := 0; index < len(discounts); index++ {
		select {
		case flagUpdate := <-ch:
			if !flagUpdate {
				s := "Error updating product"
				err := errors.New(s)
				return s, err
			}
		}
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		return "error committing queries", err
	}
	return "", nil
}

func updateDiscount(tx *sql.Tx, d map[string]string, ch chan bool) {
	//make the valuesArr to pass to the query
	valuesArr := make([]interface{}, 0)

	//pass the product_quantity, product_discount and id to the valuesArr
	valuesArr = append(valuesArr, d["product_quantity"])
	valuesArr = append(valuesArr, d["product_discount"])
	valuesArr = append(valuesArr, d["id"])

	//make query, execute query, check for errors if found send false to channel, rollback transaction and return from function. otherwise send true to the channel
	stmt := "UPDATE `discount_prices` SET `product_quantity` = ?, `product_discount` = ? WHERE `id` = ?"
	_, err := tx.ExecContext(context.Background(), stmt, valuesArr...)
	if err != nil {
		//rollback transaction send false bool to channel
		tx.Rollback()
		ch <- false
		return
	}

	//if no errors found send true to the channel
	ch <- true
}
