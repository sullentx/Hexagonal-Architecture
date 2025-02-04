package domain

import (
	entities "tienda/src/Client/Domain/entities"
)

type ClientRepository interface {
	Create(client entities.Client) error
	Modify(id int, client entities.Client) error
	Delete(id int) error
	GetClients() ([]entities.Client, error)
	Search(id int) (entities.Client, error)
}
