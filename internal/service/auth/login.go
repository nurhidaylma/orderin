package auth

import (
	"context"

	domain "github.com/nurhidaylma/orderin/internal/domain/user"
)

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if !s.hasher.Compare(user.Password, password) {
		return "", domain.ErrWrongPassword
	}

	return s.tokenGen.Generate(user)
}
