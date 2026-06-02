package repositories

import (
	"context"

	"ruach-ai/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User

	query := `
		SELECT
			id,
			email,
			password_hash,
			first_name,
			last_name,
			created_at
		FROM users
		WHERE id = $1
	`

	err := r.DB.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
