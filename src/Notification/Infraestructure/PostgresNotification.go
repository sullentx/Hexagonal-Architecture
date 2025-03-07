package infraestructure

import (
	"database/sql"
	"errors"
	entitiesC "tienda/src/client/domain/entities"
	entities "tienda/src/notification/domain/entities"
)

type PostgresNotification struct {
	db *sql.DB
}

//constructor

func NewPostgresNotificationRepository(db *sql.DB) *PostgresNotification {
	return &PostgresNotification{db: db}
}

func (r *PostgresNotification) Send(notification entities.Notification, client entitiesC.Client) error {
	_, err := r.db.Exec("INSERT INTO notifications (client_id, content) VALUES ($1, $2)", notification.ClientID, notification.Content)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresNotification) GetMessages() ([]entities.Notification, error) {
	rows, err := r.db.Query("SELECT message_id, notification_content, client_id, client_name FROM notifications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []entities.Notification
	for rows.Next() {
		var notification entities.Notification
		var clientName string
		if err := rows.Scan(&notification.ID, &notification.Content, &notification.ClientID, &clientName); err != nil {
			return nil, err
		}

		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *PostgresNotification) Search(id int) (entities.Notification, error) {
	row := r.db.QueryRow("SELECT id, client_id, content FROM notifications WHERE id = $1", id)
	var notification entities.Notification
	if err := row.Scan(&notification.ID, &notification.ClientID, &notification.Content); err != nil {
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

func (r *PostgresNotification) GetNotifications(clientID int) ([]entities.Notification, error) {
	query := "SELECT id, client_id, content FROM notifications WHERE client_id = $1"
	rows, err := r.db.Query(query, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []entities.Notification
	for rows.Next() {
		var n entities.Notification
		if err := rows.Scan(&n.ID, &n.ClientID, &n.Content); err != nil {
			return nil, err
		}
		notifications = append(notifications, n)
	}

	return notifications, nil
}
