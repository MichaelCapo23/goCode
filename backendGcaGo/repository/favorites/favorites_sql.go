package repository

import (
	"backendGcaGo/models"
	"backendGcaGo/utils"
	"database/sql"
	"errors"
)

type FavoritesRepo struct{}

func (fs FavoritesRepo) GetFavoritesRepo(db *sql.DB, favorite models.Favorite, res []models.Favorite, h map[string]string, p map[string]string) ([]models.Favorite, error) {
	//make the valuesArr to pass to the query
	valuesArr := make([]interface{}, 0)
	valuesArr = append(valuesArr, h["store_id"], p["app_user_id"])

	//make the query
	stmt := "SELECT `id`, `product_id`, `store_id` FROM `favorites` WHERE `store_id` = ? AND `app_user_id` = ? ORDER BY `created_at`"

	//execute query, pass valuesArr values to query
	rows, err := db.Query(stmt, valuesArr...)

	//check for errors
	if err != nil {
		return res, err
	}

	//scan rows set the values to favorites, append to the res
	for rows.Next() {
		err = rows.Scan(&favorite.ID, &favorite.ProductID, &favorite.StoreID)
		res = append(res, favorite)
	}

	//check for errors from scanning rows
	if err != nil {
		return []models.Favorite{}, err
	}

	//make new channel to pass to goroutines
	ch := make(chan bool)

	//loop over res, pass each row, get the img_url and product name for each row
	for i, m := range res {
		go utils.GetProductNameImg(db, res, m, i, ch)
		if flag := <-ch; !flag {
			err := errors.New("Error getting favorites")
			return []models.Favorite{}, err
		}
	}

	return res, nil
}
