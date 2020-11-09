package repository

import (
	"backendGcaGo/models"
	"database/sql"
	"fmt"
)

func (ps ProductRepo) GetProductRepo(db *sql.DB, product models.Product, res []models.Product, id string) ([]models.Product, error) {
	fmt.Println("here!!")

	return res, nil
}
