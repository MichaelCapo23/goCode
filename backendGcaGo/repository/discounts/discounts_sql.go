package repository

import (
	"backendGcaGo/models"
	"database/sql"
	"errors"
)

//DiscountRepo a struct to run on all _sql functions as common data
type DiscountRepo struct{}

//GetDiscountsRepo gets all of the valid discounts (ones that haven't ended yet) for a specific store
func (ds DiscountRepo) GetDiscountsRepo(db *sql.DB, discount models.Discount, res [][]models.Discount, h map[string]string, p map[string]string) ([][]models.Discount, error) {
	//define the slice of interfaces to pass to the prepared statement
	valuesArr := make([]interface{}, 0)

	//pass the store_id to the valuesArr
	valuesArr = append(valuesArr, h["store_id"])

	//define the query to get the valid discounts for the store that matches the store_id given in the headers
	stmt := "SELECT `id`, `product_id`, `start_date`, `end_date` FROM `discounts` WHERE (`start_date` < NOW() AND `end_date` > NOW()) AND `store_id` = ? AND `disabled` = 0 ORDER BY `created_at` DESC"

	//run the query, check for errors
	rows, err := db.Query(stmt, valuesArr...)
	if err != nil {
		return res, err
	}

	//define a new temporary slice of type models.Discount to store the rows gotten from the query
	models := []models.Discount{}

	//scan the rows from the query into the discount model, append the models
	for rows.Next() {
		err = rows.Scan(&discount.ID, &discount.ProductID, &discount.StartDate, &discount.EndDate)
		models = append(models, discount)
	}

	//check for error when scanning rows
	if err != nil {
		return res, err
	}

	//create new channel to pass to the goroutines
	ch := make(chan bool, len(models))

	//loop over the models slice
	for i, d := range models {
		//pass the db, the current row, a pointer to the response model, the channel and the current loop index
		go getDiscountPrices(db, d, &res, ch, i)

		//if the pointer value is false, there was an error getting the prices. send back error
		if flag := <-ch; !flag {
			err := errors.New("Error getting discount prices")
			return res, err
		}
	}
	//close the channel
	close(ch)
	return res, nil
}

//getDiscountPrices gets the prices, quantities and product names. (some functions like this will he held in the utils.go file. decided to keep this here for readability)
func getDiscountPrices(db *sql.DB, d models.Discount, res *[][]models.Discount, ch chan bool, i int) {
	//make a new slice to manage and build the current slice being made
	currentSlice := make([]models.Discount, 0)

	//get the product name of the current product
	stmt := "SELECT `product_name` FROM `products_master_list` WHERE `id` = ?"
	err := db.QueryRow(stmt, d.ProductID).Scan(&d.ProductName)
	if err != nil {
		ch <- false
	}

	//get the discount information for the current discount id
	stmt = "SELECT `id`, `discount_id`, `product_quantity`, `product_discount` FROM `discount_prices` WHERE `discount_id` = ?"
	rows, err2 := db.Query(stmt, d.ID)
	if err2 != nil {
		ch <- false
	}

	//scan the rows assign values of the current row to d (current row passed from models struct) and append it to my currentSlice slice of models.Discount
	for rows.Next() {
		err = rows.Scan(&d.ID, &d.DiscountID, &d.ProductQuantity, &d.ProductDiscount)
		currentSlice = append(currentSlice, d)
	}

	//check the scan for errors
	if err != nil {
		ch <- false
	}
	*res = append(*res, currentSlice)
	ch <- true
}
