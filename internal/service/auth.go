package service

import (
	"time"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/fishmanDK/miet_project/internal/storage"
)

type AuthService struct{
	storage *storage.Storage
}

func newAuthService(storage *storage.Storage) *AuthService{
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) Authentication(user core.Client) (core.Tokens, error){
	var tokens core.Tokens

	res, err := s.storage.Auth.Authentication(user)
	if err != nil{
		return tokens, err
	}

	tokens.Access_token, err = CreateAccessToken(res.Id, res.Email, res.Role)
	if err != nil {
		return tokens, err
	}

	tokens.Refresh_token, err = CreateRefreshToken()
	if err != nil {
		return tokens, err
	}

	session := core.Session{
		Refresh_token: tokens.Refresh_token,
		ExpiresAt:     time.Now().Add(refresh_tokenTtl).UTC(),
	}

	err = s.storage.Auth.CreateSession(res.Id, session)
	return tokens, err
}

func (s *AuthService) CreateUser(newUser core.Client) (int, error){
	id, err := s.storage.Auth.CreateUser(newUser)
	return id, err
}