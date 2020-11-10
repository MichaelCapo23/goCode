package repository

import (
	"backendGcaGo/models"
	"database/sql"
)

//GetProductRepo repo function for getting a single product by an id in the route
func (ps ProductRepo) GetProductRepo(db *sql.DB, product models.Product, id string) (models.Product, error) {
	stmt := "SELECT DISTINCT `id`, `product_name`, `description`, `category_id`, `subcategory_id`, `status`, `manufacturer`, `recommended` FROM `products_master_list` WHERE `id` = ?"
	rows, err := db.Query(stmt, id)

	for rows.Next() {
		var manufacturer sql.NullString
		err = rows.Scan(&product.ID, &product.ProductName, &product.Description, &product.CategoryID, &product.SubcategoryID, &product.Status, &manufacturer, &product.Recommended)
		product.Manufacturer = manufacturer.String
	}

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
