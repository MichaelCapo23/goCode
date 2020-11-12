package repository

import (
	"backendGcaGo/driver"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

//AddDiscountRepo adds/updates discounts based of the id in the "discount" slice of maps
func (ds DiscountRepo) AddDiscountRepo(db *sql.DB, h map[string]string, p map[string]string) (string, error) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		driver.LogFatal(err)
		fmt.Println("here daddy1")
		return "", err
	}

	ch := make(chan string)

	// //ge the id of the
	// go func() {

	// }()

	//make a new discounts variable and parse the slice of maps set the value of the new discounts variable
	discounts := make([]map[string]string, 0)
	json.Unmarshal([]byte(p["discounts"]), &discounts)

	//make 2 new points one for the new discounts another for the ones being updated
	chNew := make(chan bool, len(discounts))
	chUpdate := make(chan bool, len(discounts))

	//start to update the/wait till the channel is filled to create new discounts

	for _, d := range discounts {
		if d["id"] == "-1" {
			go addNewDiscount(tx, d, chNew, ch)
		} else {
			go updateDiscount(tx, d, chUpdate)
		}
	}

	//defer then close channels
	defer close(chNew)
	defer close(chUpdate)

	for index := 0; index < len(discounts); index++ {
		select {
		case flagNew := <-chNew:
			if !flagNew {
				s := "Error adding product"
				err := errors.New(s)
				return s, err
			}
		case flagUpdate := <-chUpdate:
			if !flagUpdate {
				s := "Error updating product"
				err := errors.New(s)
				return s, err
			}
		}
	}

	return "", nil
}

func addNewDiscount(tx *sql.Tx, d map[string]string, chNew chan bool, ch chan string) {
	//set the channel that got the id of the parent discount to the value of dID
	dID := <-ch

	//define a valuesArr to pass to the query
	valuesArr := make([]interface{}, 0)

	//add values to the slice to use in the query
	valuesArr = append(valuesArr, dID)
	valuesArr = append(valuesArr, d["product_quantity"])
	valuesArr = append(valuesArr, d["product_discount"])

	//build the query
	stmt := "INSERT INTO `discount_prices` (`discount_id`, `product_quantity`, `product_discount`) VALUES (?,?,?)"

	//call the query, check for errors. if error found send false to channel return out, otherwise send true
	_, err := tx.ExecContext(context.Background(), stmt, valuesArr...)
	if err != nil {
		tx.Rollback()
		chNew <- false
		return
	}
	chNew <- true
}

func updateDiscount(tx *sql.Tx, d map[string]string, chUpdate chan bool) {
	fmt.Println("update" + d["product_discount"])
	chUpdate <- true
}
