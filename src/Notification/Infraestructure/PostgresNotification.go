package infraestructure

import (
	"database/sql"
	"errors"
	entities "tienda/src/Notification/Domain/Entities"
)

type PostgresNotification struct {
	db *sql.DB
}

//constructor

func NewPostgresNotificationRepository(db *sql.DB) *PostgresNotification {
	return &PostgresNotification{db: db}
}

func (r PostgresNotification) Send(notification entities.Notification) error {
	_, err := r.db.Exec("INSERT INTO notifications (content) VALUES ($1)", notification.Content)
	return err
}

func (r PostgresNotification) GetMessages() ([]entities.Notification, error) {
	rows, err := r.db.Query("SELECT id, content FROM notifications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []entities.Notification
	for rows.Next() {
		var notification entities.Notification
		if err := rows.Scan(&notification.ID, &notification.Content); err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}

func (r PostgresNotification) Search(id int) (entities.Notification, error) {
	row := r.db.QueryRow("SELECT id, content FROM notifications WHERE id = $1", id)
	var notification entities.Notification
	if err := row.Scan(&notification.ID, &notification.Content); err != nil {
		return entities.Notification{}, err
	}
	return notification, nil
}

func (r PostgresNotification) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM notifications WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("notification not found")
	}
	return nil
}

func (r PostgresNotification) ModifyMessage(id int, content string) error {
	result, err := r.db.Exec("UPDATE notifications SET content = $1 WHERE id = $2", content, id)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("notification not found")
	}
	return nil
}
