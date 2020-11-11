package repository

import (
	"backendGcaGo/driver"
	"backendGcaGo/models"
	"context"
	"database/sql"
)

//DiscountRepo a struct to run on all _sql functions as common data
type DiscountRepo struct{}

//GetDiscountsRepo gets all of the discounts for a specific store, that end at any point past when the service is called
func (ds DiscountRepo) GetDiscountsRepo(db *sql.DB, discount models.Discount, res []models.Discount, h map[string]string, p map[string]string) ([]models.Discount, error) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		driver.LogFatal(err)
	}

	//define the slice of interfaces to pass to the prepared statement
	valuesArr := make([]interface{}, 0)
	valuesArr = append(valuesArr, h["store_id"])

	stmt := "SELECT `id`, `product_id`, `start_date`, `end_date`, `disabled` FROM `discounts` WHERE (`start_date` < NOW() AND `end_date` > NOW()) AND `store_id` = ? AND `disabled` = 0 ORDER BY `created_at` DESC"

	tx.ExecContext(ctx, stmt, valuesArr...)
	return res, nil
}
