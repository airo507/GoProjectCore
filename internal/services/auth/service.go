package auth

import "context"

type AuthService struct {
}

func NewRegistrationService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Authorize(ctx context.Context) error {
	return nil
}
