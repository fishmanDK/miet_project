package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	refresh_tokenTtl = time.Hour * 24 * 30
	access_tokenTtl  = time.Minute * 15
	signInKey        = "@(#tf53$*#$(RHfverib}#Rfrte)"
	salt             = "lsd2#tfv%2"
)

type Claims struct {
	jwt.StandardClaims
	Id    int
	Email string
	Role  string
}

type ParseDataUser struct {
	ID    int
	Email string
	Role  string
}

func CreateAccessToken(id int, email, role string) (string, error) {

	acssesToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(access_tokenTtl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:    id,
		Email: email,
		Role:  role,
	})
	return acssesToken.SignedString([]byte(signInKey))
}

func CreateRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	_, err := r.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

func (s *AuthService) ParseToken(accessToken string) (*ParseDataUser, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s: %w", errors.New("token verification error"))
		}
		return []byte(signInKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("failed to parse token claims")
	}

	if !ok {
		return nil, err
	}

	res := ParseDataUser{
		ID:    claims.Id,
		Email: claims.Email,
		Role:  claims.Role,
	}

	return &res, nil
}