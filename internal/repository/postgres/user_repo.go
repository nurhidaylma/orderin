package postgres

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	domain "github.com/nurhidaylma/orderin/internal/domain/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (id, name, email, password_hash, role, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		user.CreatedAt,
	)
	if err != nil {
		if isUniqueViolationError(err) {
			return domain.ErrEmailAlreadyExists
		}
	}

	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	return nil, nil
}

func isUniqueViolationError(err error) bool {
	if e, ok := err.(*pq.Error); ok {
		return e.Code == "23505"
	}
	return false
}
