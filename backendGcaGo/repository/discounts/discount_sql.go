package repository

import (
	"backendGcaGo/models"
	"database/sql"
)

// GetDiscountRepo gets a single discount based of id in url
func (ds DiscountRepo) GetDiscountRepo(db *sql.DB, discount models.Discount, h map[string]string, id string) ([]models.Discount, error) {
	stmt := "SELECT `d`.`id`, `d`.`product_id`, `d`.`start_date`, `d`.`end_date`, `pml`.`product_name`, `dp`.`id`, `dp`.`discount_id`, `dp`.`product_quantity`, `dp`.`product_discount` FROM `discounts` AS `d` INNER JOIN `products_master_list` AS `pml` ON (`pml`.`id` = `d`.product_id) INNER JOIN `discount_prices` AS `dp` ON (`dp`.`discount_id` = `d`.id) WHERE (`start_date` < NOW() AND `end_date` > NOW()) AND `d`.`store_id` = ? AND `d`.`disabled` = 0 AND `d`.`id` = ?"

	//make and append the values to the slice to use in the query
	valuesArr, res := make([]interface{}, 0), []models.Discount{}
	valuesArr = append(valuesArr, h["store_id"])
	valuesArr = append(valuesArr, id)

	//run the query check for errors
	rows, err := db.Query(stmt, valuesArr...)
	if err != nil {
		return res, err
	}

	//scan rows assign values to discount data models, append to res slice of type models.Discount
	for rows.Next() {
		err = rows.Scan(&discount.ID, &discount.ProductID, &discount.StartDate, &discount.EndDate, &discount.ProductName, &discount.ID, &discount.DiscountID, &discount.ProductQuantity, &discount.ProductDiscount)
		res = append(res, discount)
	}

	//check for errors
	if err != nil {
		return res, err
	}

	//return data
	return res, nil
}
