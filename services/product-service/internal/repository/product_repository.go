package repository

import (
	"context"
	"product-service/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func (r *ProductRepository) Create(
	ctx context.Context,
	product *model.Product,
) error {
	_, err := r.db.Exec(
		ctx,
		`
		INSERT INTO products
		(id, name, description, price)
		VALUE ($1, $2, $3, $4)
		`,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
	)

	return err
}

func (r *ProductRepository) GetAll(
	ctx context.Context,
) ([]model.Product, error) {
	rows, err := r.db.Query(
		ctx,
		`
		SELECT id, name, description, price, created_at, updated_at
		FROM products
		ORDER BY created_at DESC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var product model.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, rows.Err()
}