package infraestructure

//LSP Liskov Substitution Principle (Principio de sustitución de Liskov)
//Se puede sustituir un tipo por otro si cumple con la misma interfaz
import (
	"database/sql"
	"errors"
	entities "tienda/src/Products/Domain-negocio/entities"
)

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) Save(product entities.Product) error {
	_, err := r.db.Exec("INSERT INTO products (name, price, quantity) VALUES ($1, $2, $3)", product.Name, product.Price, product.Quantity)
	return err
}

func (r *PostgresProductRepository) GetAll() ([]entities.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price, quantity FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *PostgresProductRepository) GetOne(id int) (entities.Product, error) {
	row := r.db.QueryRow("SELECT id,name, price, quantity FROM products WHERE id = $1", id)
	var product entities.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity); err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

// duda esto puede retornar un texto?
func (r *PostgresProductRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM products WHERE id =$1", id)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("producto no encontrado")
	}
	return nil
}

func (r *PostgresProductRepository) Put(id int, product entities.Product) error {
	result, err := r.db.Exec("UPDATE products SET name = $1, price = $2 , quantity = $3 WHERE id = $4", product.Name, product.Price, product.Quantity, id)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errors.New("producto no encontrado")
	}
	return nil
}

func (r *PostgresProductRepository) GetNewProducts(lastProductID int) ([]entities.Product, int, error) {
	rows, err := r.db.Query("SELECT id, name, price, quantity FROM products WHERE id > $1", lastProductID)
	if err != nil {
		return nil, lastProductID, err
	}
	defer rows.Close()

	var products []entities.Product
	var latestID int
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity); err != nil {
			return nil, lastProductID, err
		}
		products = append(products, product)
		latestID = product.ID
	}
	return products, latestID, nil
}
