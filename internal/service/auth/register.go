package auth

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/nurhidaylma/orderin/internal/domain/user"
)

func (s *authService) Register(
	ctx context.Context,
	name, email, password string,
	role domain.Role,
) (*domain.User, error) {

	if role != domain.RoleUser && role != domain.RoleMerchant {
		return nil, domain.ErrInvalidRole
	}

	hash, err := s.hasher.Hash(password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:       uuid.NewString(),
		Name:     name,
		Email:    email,
		Password: hash,
		Role:     role,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
