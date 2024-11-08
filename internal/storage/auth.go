package storage

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/jmoiron/sqlx"
)

type AuthStorage struct{
	db *sqlx.DB
}

func newAuthStorage(db *sqlx.DB) *AuthStorage{
	return &AuthStorage{
		db: db,
	}
}

func (s *AuthStorage) CreateUser(newUser core.Client) (int, error) {
	query := sq.Insert("users").
	Columns("email", "password", "role").
	Values(newUser.Email, newUser.Password, "user").
	Suffix("RETURNING \"id\"").
	RunWith(s.db).
    PlaceholderFormat(sq.Dollar)

	var id int
	err := query.QueryRow().Scan(&id)
	if err != nil{
		return 0, err
	}

	return id, nil
}

func (s *AuthStorage) Authentication(user core.Client) (core.AuthResult, error){
	query := sq.Select("id", "email", "role").	
	From("users").
	Where(sq.Eq{"email": user.Email}).
	RunWith(s.db).
    PlaceholderFormat(sq.Dollar)

	var res core.AuthResult

	err := query.QueryRow().Scan(&res.Id, &res.Email, &res.Role)
	if err != nil{
		return core.AuthResult{}, err
	}

	return res, nil
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