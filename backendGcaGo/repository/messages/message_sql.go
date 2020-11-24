package repository

import (
	"backendGcaGo/models"
	"database/sql"
)

// GetMessageRepo gets a single message based off of the id passed
func (ms MessagesRepo) GetMessageRepo(db *sql.DB, message models.Message, h map[string]string, id string) (models.Message, error) {
	//create valuesArr to pass to the query
	valuesArr := make([]interface{}, 0)
	valuesArr = append(valuesArr, h["store_id"], id)

	//build the query
	stmt := "SELECT `m`.`id`, `m`.`type`, `m`.`app_user_id`, `m`.`message`, `m`.`resolved`, `m`.`created_at`, `au`.`first_name`, `au`.`last_name`, `au`.`uuid` FROM `messages` AS `m` JOIN `app_users` AS `au` ON (`m`.`app_user_id` = `au`.`id`) WHERE `store_id` = ? AND `m`.`id` = ?"
	err := db.QueryRow(stmt, valuesArr...).Scan(&message.ID, &message.Type, &message.AppUserID, &message.Message, &message.Resolved, &message.CreatedAt, &message.FirstName, &message.LastName, &message.UUID)
	if err != nil {
		return models.Message{}, err
	}
	return message, nil
}
