package repository

import (
	"backendGcaGo/models"
	"database/sql"
	"errors"
)

type MessagesRepo struct{}

//GetMessagesRepo gets all of the messages for each store. filtered by the `type` passed in the parameters
func (ms MessagesRepo) GetMessagesRepo(db *sql.DB, message models.Message, res []models.Message, p map[string]string, h map[string]string) ([]models.Message, error) {
	//create the valuesArr to pass to the query
	valuesArr := make([]interface{}, 0)
	valuesArr = append(valuesArr, h["store_id"], p["type"])

	//build the query
	stmt := "SELECT `id`, `type`, `app_user_id`, `message`, `resolved`, `created_at` FROM `messages` WHERE `store_id` = ? AND `type` = ? ORDER BY `created_at` DESC"
	rows, err := db.Query(stmt, valuesArr...)
	if err != nil {
		return []models.Message{}, err
	}

	//loop over the rows, set values to message. Append message to res
	for rows.Next() {
		err = rows.Scan(&message.ID, &message.Type, &message.AppUserID, &message.Message, &message.Resolved, &message.CreatedAt)
		res = append(res, message)
	}

	//check for errors while scanning
	if err != nil {
		return []models.Message{}, err
	}

	//next section gets the app_user information (first name, last name, and uuid)

	//make channel to pass to goroutines
	ch := make(chan bool)

	for i, r := range res {
		go getMessageUserInfo(db, res, r, i, ch)

		select {
		case flag := <-ch:
			if !flag {
				return []models.Message{}, errors.New("error getting user information")
			}
		}
	}

	return res, nil
}

func getMessageUserInfo(db *sql.DB, res []models.Message, r models.Message, index int, ch chan bool) {
	stmt := "SELECT `first_name`, `last_name`, `uuid` FROM `app_users` WHERE `id` = ?"
	err := db.QueryRow(stmt, r.AppUserID).Scan(&r.FirstName, &r.LastName, &r.UUID)

	if err != nil {
		ch <- false
	}

	res[index] = r
	ch <- true
}
