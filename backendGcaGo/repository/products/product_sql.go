package repository

import (
	"backendGcaGo/models"
	"database/sql"
)

//GetProductRepo repo function for getting a single product by an id in the route
func (ps ProductRepo) GetProductRepo(db *sql.DB, product models.Product, id string) (models.Product, error) {
	//build the query
	stmt := "SELECT DISTINCT `id`, `product_name`, `description`, `category_id`, `subcategory_id`, `status`, `manufacturer`, `recommended` FROM `products_master_list` WHERE `id` = ?"

	//run the query check for errors
	rows, err := db.Query(stmt, id)
	if err != err {
		return models.Product{}, err
	}

	//scan the rows
	for rows.Next() {
		//make a variable of type sql.NullString
		var manufacturer sql.NullString
		err = rows.Scan(&product.ID, &product.ProductName, &product.Description, &product.CategoryID, &product.SubcategoryID, &product.Status, &manufacturer, &product.Recommended)
		//add manufacturer.String to the product.Manufacturer to get an empty string if null found or the value if a value was found
		product.Manufacturer = manufacturer.String
	}

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
