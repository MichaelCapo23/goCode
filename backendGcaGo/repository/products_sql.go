package repository

import (
	"backendGcaGo/models"
	"database/sql"
)

type ProductRepo struct{}

func (p ProductRepo) GetProductsRepo(db *sql.DB, product models.Product, res []models.Product) ([]models.Product, error) {
	rows, err := db.Query("SELECT product_name, description FROM products_master_list")
	if err != nil {
		return []models.Product{}, err
	}

	for rows.Next() {
		err = rows.Scan(&product.Name, &product.Description)
		res = append(res, product)
	}

	if err != nil {
		return []models.Product{}, err
	}

	return res, nil
}
