package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PasswordRepository struct {
	db *pgxpool.Pool
}

func NewPasswordRepository(db *pgxpool.Pool) *PasswordRepository {
	return &PasswordRepository{
		db: db,
	}
}

func (r *PasswordRepository) CreateResetToken(
	ctx context.Context,
	userID string,
	hashedToken string,
	expiresAt time.Time,
) error {
	_, err := r.db.Exec(
		ctx,
		`
		INSERT INTO password_reset_tokens (
			user_id,
			token_hash,
			expires_at
		)
		VALUES ($1, $2, $3)
		`,
		userID,
		hashedToken,
		expiresAt,
	)

	return err
}