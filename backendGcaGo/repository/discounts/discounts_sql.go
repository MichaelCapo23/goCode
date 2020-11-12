package repository

import (
	"backendGcaGo/driver"
	"backendGcaGo/models"
	"database/sql"
	"errors"
)

//DiscountRepo a struct to run on all _sql functions as common data
type DiscountRepo struct{}

//GetDiscountsRepo gets all of the discounts for a specific store, that end at any point past when the service is called
func (ds DiscountRepo) GetDiscountsRepo(db *sql.DB, discount models.Discount, res [][]models.Discount, h map[string]string, p map[string]string) ([][]models.Discount, error) {
	//define the slice of interfaces to pass to the prepared statement
	valuesArr := make([]interface{}, 0)
	valuesArr = append(valuesArr, h["store_id"])

	stmt := "SELECT `id`, `product_id`, `start_date`, `end_date` FROM `discounts` WHERE (`start_date` < NOW() AND `end_date` > NOW()) AND `store_id` = ? AND `disabled` = 0 ORDER BY `created_at` DESC"

	rows, err := db.Query(stmt, valuesArr...)
	if err != nil {
		driver.LogFatal(err)
		return res, err
	}

	models := []models.Discount{}

	for rows.Next() {
		err = rows.Scan(&discount.ID, &discount.ProductID, &discount.StartDate, &discount.EndDate)
		models = append(models, discount)
	}

	if err != nil {
		driver.LogFatal(err)
		return res, err
	}

	ch := make(chan bool, len(models))

	for i, d := range models {
		go getDiscountPrices(db, d, &res, ch, i)
		if flag := <-ch; !flag {
			err := errors.New("Error adding new prices")
			return res, err
		}
	}
	close(ch)
	return res, nil
}

func getDiscountPrices(db *sql.DB, d models.Discount, res *[][]models.Discount, ch chan bool, i int) {
	//make a new slice to manage and build the current slice being made
	currentSlice := make([]models.Discount, 0)

	//get the product name of the current product
	stmt := "SELECT `product_name` FROM `products_master_list` WHERE `id` = ?"
	err := db.QueryRow(stmt, d.ProductID).Scan(&d.ProductName)
	if err != nil {
		driver.LogFatal(err)
		ch <- false
	}

	//get the discount information for the current discount id
	stmt = "SELECT `id`, `discount_id`, `product_quantity`, `product_discount` FROM `discount_prices` WHERE `discount_id` = ?"
	rows, err2 := db.Query(stmt, d.ID)
	if err2 != nil {
		driver.LogFatal(err)
		ch <- false
	}

	for rows.Next() {
		err = rows.Scan(&d.ID, &d.DiscountID, &d.ProductQuantity, &d.ProductDiscount)
		currentSlice = append(currentSlice, d)
	}

	//check the scan for errors
	if err != nil {
		driver.LogFatal(err)
		ch <- false
	}
	*res = append(*res, currentSlice)
	ch <- true
}
