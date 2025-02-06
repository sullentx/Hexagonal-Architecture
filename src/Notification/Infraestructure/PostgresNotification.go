package infraestructure

import (
	"database/sql"
	"errors"
	entitiesC "tienda/src/Client/Domain/entities"
	entities "tienda/src/Notification/Domain/Entities"
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

	// Simular el envío de la notificación por correo electrónico

	return nil
}

func (r *PostgresNotification) GetMessages() ([]entities.Notification, error) {
	rows, err := r.db.Query("SELECT id, client_id, content FROM notifications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []entities.Notification
	for rows.Next() {
		var notification entities.Notification
		if err := rows.Scan(&notification.ID, &notification.ClientID, &notification.Content); err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
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
