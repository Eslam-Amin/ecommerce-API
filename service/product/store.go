package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Eslam-Amin/ecommerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateProduct(product types.Product) error {
	_, err := s.db.Exec(`
	INSERT INTO products 
	(name, description, image, price, quantity)
	VALUES (?,?,?,?,?)
	`, product.Name, product.Description, product.Image, product.Price, product.Quantity)
	return err
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query(`
	SELECT id, name, description, image, price, quantity, createdAt
	FROM products
`)
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) GetProductsByIds(productIds []int) ([]types.Product, error) {
	placeholders := strings.Repeat(",?", len(productIds)-1)
	query := fmt.Sprintf(`
	SELECT id, name, description, image, price, quantity, createdAt
	FROM products
	where id IN (?%s)
	`, placeholders)

	// Convert ProductIds to []interface{}
	args := make([]interface{}, len(productIds))
	for i, v := range productIds {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}
