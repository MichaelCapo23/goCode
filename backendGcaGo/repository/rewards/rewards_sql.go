package repository

import (
	"backendGcaGo/models"
	"database/sql"
)

type RewardsRepo struct{}

//GetRewardsRepo gets the rewards for a single user
func (rs RewardsRepo) GetRewardsRepo(db *sql.DB, reward models.Reward, res []models.Reward, h map[string]string, id string) ([]models.Reward, error) {
	//make the valuesArr, append items to pass to the query
	valuesArr := make([]interface{}, 0)
	valuesArr = append(valuesArr, h["store_id"], id)

	//make the query to get all rewards for a specific person
	stmt := "SELECT `id`, `point_count`, `claimed`, `claimed_date`, `created_at` FROM `rewards` WHERE `store_id` = ? AND `app_user_id` = ? ORDER BY `created_at` DESC"

	//run query, check for errors
	rows, err := db.Query(stmt, valuesArr...)
	if err != nil {
		return res, err
	}

	//scan rows pass values to reward (type models.Reward) and append it to res (slice of type models.Reward)
	for rows.Next() {
		err = rows.Scan(&reward.ID, &reward.PointCount, &reward.Claimed, &reward.ClaimedDate, &reward.CreatedAt)
		res = append(res, reward)
	}

	//check for errors
	if err != nil {
		return res, err
	}
	return res, nil
}
