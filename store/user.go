package store

import (
	"context"

	"github.com/aksbuzz/mood-journal/api"
)

func (s *Store) CreateUser(ctx context.Context, createUser *api.User) (*api.User, error) {
	query := `
	INSERT INTO users (
		username,
		password_hash
	)
	VALUES (?, ?)
	RETURNING id, created_at, updated_at, username
	`
	var user api.User
	if err := s.db.QueryRowContext(
		ctx, query, createUser.Username, createUser.PasswordHash,
	).Scan(
		&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Username,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Store) FindUser(ctx context.Context, findUser *api.FindUser) (*api.User, error) {
	list, err := s.ListUsers(ctx, findUser)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	user := list[0]
	return user, nil
}

func (s *Store) ListUsers(ctx context.Context, findUser *api.FindUser) ([]*api.User, error) {
	where, args := []string{"1 = 1"}, []any{}

	if v := findUser.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}
	if v := findUser.UserName; v != nil {
		where, args = append(where, "username = ?"), append(args, *v)
	}
	if v := findUser.DisplayName; v != nil {
		where, args = append(where, "display_name = ?"), append(args, *v)
	}
	if v := findUser.Email; v != nil {
		where, args = append(where, "email = ?"), append(args, *v)
	}

	query := `SELECT
		id, 
		created_at, 
		updated_at, 
		username, 
		display_name, 
		email, 
		password_hash, 
		avatar_url 
	FROM users
	WHERE ` + joinStrings(where, " AND ") + `
	ORDER BY created_at DESC
	`

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*api.User, 0)
	for rows.Next() {
		var user api.User
		if err := rows.Scan(
			&user.ID,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Username,
			&user.DisplayName,
			&user.Email,
			&user.PasswordHash,
			&user.AvatarURL,
		); err != nil {
			return nil, err
		}
		list = append(list, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
