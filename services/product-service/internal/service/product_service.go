package service

import (
	"context"
	"errors"
	"product-service/internal/model"
	"product-service/internal/repository"

	"github.com/google/uuid"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(
	repo *repository.ProductRepository,
) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) Create(
	ctx context.Context,
	name string,
	description string,
	price float64,
) (*model.Product, error) {
	if name == "" {
		return nil, errors.New("product name is required")
	}

	if price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}

	product := &model.Product{
		ID: uuid.New().String(),
		Name: name,
		Description: description,
		Price: price,
	}

	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}