package repository

import (
	"auth-service/internal/model"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	user *model.User,
) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO users (id, email, password_hash, role)
		VALUES ($1, $2, $3, $4)`,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.Role,
	)

	return  err
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*model.User, error) {
	user := &model.User{}

	err := r.db.QueryRow(
		ctx,
		`
		SELECT id, email, password_hash, role, created_at
		FROM users
		WHERE email = $1
		`,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil

}