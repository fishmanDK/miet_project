package storage

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/doug-martin/goqu/v9"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/jmoiron/sqlx"
)

type AuthStorage struct {
	db *sqlx.DB
}

func NewAuthStorage(db *sqlx.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}

func (s *AuthStorage) CreateUser(newUser core.Client) (int, error) {
	emptyUser := core.Client{}
	if newUser == emptyUser {
		return 0, errors.New("empty new user")
	}
	newUser.Role = "user"
	qu, _, err := goqu.Insert("users").Rows(newUser).Returning("id").ToSQL()
	if err != nil {
		return 0, err
	}

	var id int
	if err := s.db.QueryRowxContext(context.Background(), qu).Scan(&id); err != nil {
		return 0, fmt.Errorf("create new user error during execute query: %w", err)
	}

	return id, nil
}

func toDatasetAuth(user *core.Client) *goqu.SelectDataset {
	selectDataset := goqu.From("users").Select("id", "email", "role")

	if user == nil {
		return selectDataset
	}

	if user.Email != "" {
		selectDataset = selectDataset.Where(goqu.Ex{"email": user.Email})
	}

	return selectDataset
}

func (s *AuthStorage) Authentication(user core.Client) (core.AuthResult, error) {
	query, _, err := toDatasetAuth(&user).ToSQL()
	if err != nil {
		return core.AuthResult{}, err
	}

	var res []core.AuthResult
	if err = s.db.SelectContext(context.Background(), &res, query); err != nil {
		return core.AuthResult{}, fmt.Errorf("select user error during execute query: %w", err)
	}

	return res[0], nil
}

func (s *AuthStorage) CreateSession(userId int, session core.Session) error {
	query, args, err := sq.Insert("jwt").
		Columns("user_id", "refresh_token", "expiresAt").
		Values(userId, session.Refresh_token, session.ExpiresAt).
		Suffix("ON CONFLICT (user_id) DO UPDATE SET refresh_token = EXCLUDED.refresh_token, expiresAt = EXCLUDED.expiresAt").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = s.db.Exec(query, args...)
	return err
}
