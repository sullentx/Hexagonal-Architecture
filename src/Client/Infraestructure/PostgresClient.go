package infraestructure

import (
	"database/sql"
	"errors"
	"tienda/src/Client/Domain/entities"
)

type PostgresClientRepository struct {
	db *sql.DB
}

func NewPostgresClientRepository(db *sql.DB) *PostgresClientRepository {
	return &PostgresClientRepository{db: db}
}

func (r *PostgresClientRepository) Create(client entities.Client) error {
	_, err := r.db.Exec("INSERT INTO clients (name, last_name, email, password) VALUES ($1, $2, $3, $4)", client.Name, client.LastName, client.Email, client.Password)
	return err
}

func (r *PostgresClientRepository) Search(id int) (entities.Client, error) {
	row := r.db.QueryRow("SELECT id, name, last_name, email, password FROM clients WHERE id = $1", id)
	var client entities.Client
	err := row.Scan(&client.ID, &client.Name, &client.LastName, &client.Email, &client.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Client{}, errors.New("client not found")
		}
		return entities.Client{}, err
	}
	return client, nil
}

func (r *PostgresClientRepository) GetClients() ([]entities.Client, error) {
	rows, err := r.db.Query("SELECT id, name, last_name, email, password FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []entities.Client
	for rows.Next() {
		var client entities.Client
		if err := rows.Scan(&client.ID, &client.Name, &client.LastName, &client.Email, &client.Password); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func (r *PostgresClientRepository) Modify(id int, client entities.Client) error {
	_, err := r.db.Exec("UPDATE clients SET name = $1, last_name = $2, email = $3, password = $4 WHERE id = $5", client.Name, client.LastName, client.Email, client.Password, id)
	return err
}

func (r *PostgresClientRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM clients WHERE id = $1", id)
	return err
}
