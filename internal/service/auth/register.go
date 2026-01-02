package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurhidaylma/orderin/internal/domain/user"
)

func (s *authService) Register(
	ctx context.Context,
	name, email, password string,
	role user.Role,
) (*user.User, error) {

	if role != user.RoleUser && role != user.RoleMerchant {
		return nil, user.ErrInvalidRole
	}

	hash, err := s.hasher.Hash(password)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		ID:       generateUUID(),
		Name:     name,
		Email:    email,
		Password: hash,
		Role:     role,
	}

	if err := s.userRepo.Create(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}

func generateUUID() string {
	return uuid.New().String()
}
