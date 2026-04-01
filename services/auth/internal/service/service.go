package service

import (
	"errors"
)

var ErrUnauthorized = errors.New("unauthorized")

type AuthService struct {
	validUsername string
	validPassword string
	validToken    string
}

type VerifyResult struct {
	Valid   bool   `json:"valid"`
	Subject string `json:"subject,omitempty"`
	Error   string `json:"error,omitempty"`
}

func New() *AuthService {
	return &AuthService{
		validUsername: "student",
		validPassword: "student",
		validToken:    "demo-token",
	}
}

func (s *AuthService) Login(username, password string) (string, bool) {
	if username == s.validUsername && password == s.validPassword {
		return s.validToken, true
	}
	return "", false
}

func (s *AuthService) Verify(token string) (string, error) {
	if token == s.validToken {
		return "student", nil
	}
	return "", ErrUnauthorized
}

func (s *AuthService) VerifyHTTP(token string) VerifyResult {
	subject, err := s.Verify(token)
	if err != nil {
		return VerifyResult{
			Valid: false,
			Error: "unauthorized",
		}
	}

	return VerifyResult{
		Valid:   true,
		Subject: subject,
	}
}
