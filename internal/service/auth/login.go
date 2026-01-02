package auth

import (
	"context"

	"github.com/nurhidaylma/orderin/internal/domain/user"
)

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	u, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if !s.hasher.Compare(u.Password, password) {
		return "", user.ErrWrongPassword
	}

	return s.tokenGen.Generate(u)
}
