package auth

import (
	"context"

	"github.com/nurhidaylma/orderin/internal/domain/user"
)

type authService struct {
	userRepo user.Repository
	hasher   PasswordHasher
	tokenGen TokenGenerator
}

func NewService(
	userRepo user.Repository,
	hasher PasswordHasher,
	tokenGen TokenGenerator,
) Service {
	return &authService{
		userRepo: userRepo,
		hasher:   hasher,
		tokenGen: tokenGen,
	}
}

type Service interface {
	Register(ctx context.Context, name, email, password string, role user.Role) (*user.User, error)
	Login(ctx context.Context, email, password string) (string, error) // returns JWT
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) bool
}

type TokenGenerator interface {
	Generate(user *user.User) (string, error)
}
