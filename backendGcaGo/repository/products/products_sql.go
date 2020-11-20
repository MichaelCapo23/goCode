package repository

import (
	"backendGcaGo/models"
	"database/sql"
)

type ProductRepo struct{}

//GetProductsRepo repo function for getting a single product by an id in the route
func (ps ProductRepo) GetProductsRepo(db *sql.DB, product models.Product, res []models.Product, h map[string]string, p map[string]string) ([]models.Product, error) {
	//make the query params array
	valuesArr := make([]interface{}, 0)

	//build the base of the query
	stmt := "SELECT DISTINCT `id`, `product_name`, `description`, `category_id`, `subcategory_id`, `status`, `manufacturer`, `recommended` FROM `products_master_list` "

	//define the default where clause that grabs all of the products with a store of -1 (masterList)
	whereClause := "WHERE `store_id` = -1 ORDER BY `created_at` DESC"
	stmt += whereClause

	//check the type parameter, append the correct where clause to query, add the correct values to valuesArr to pass to the query (every item used by the store_id given)
	if p["type"] == "all" {
		whereClause = "WHERE `store_id` = ? ORDER BY `created_at` DESC"
		stmt += whereClause
		valuesArr = append(valuesArr, h["store_id"])
	}

	//run the query and check for errors
	rows, err := db.Query(stmt, valuesArr...)
	if err != nil {
		return []models.Product{}, err
	}

	//close the rows
	defer rows.Close()

	//loop over rows, scan, append to return object
	for rows.Next() {
		//define a NullString value, incase the value at that column is a null value
		var manufacturer sql.NullString
		err = rows.Scan(&product.ID, &product.ProductName, &product.Description, &product.CategoryID, &product.SubcategoryID, &product.Status, &manufacturer, &product.Recommended)
		//assign the NullString's String value to current product
		product.Manufacturer = manufacturer.String
		res = append(res, product)
	}

	//check the scan for errors
	if err != nil {
		return []models.Product{}, err
	}

	return res, nil
}
