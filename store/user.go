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
		return nil, FormatError(err)
	}
	s.userCache.Set(&user)
	return &user, nil
}

func (s *Store) UpdateUser(ctx context.Context, updateUser *api.UpdateUser) (*api.User, error) {
	set, args := []string{}, []any{}
	if v := updateUser.Username; v != nil {
		set, args = append(set, "username = ?"), append(args, *v)
	}
	if v := updateUser.UpdatedAt; v != nil {
		set, args = append(set, "updated_at = ?"), append(args, *v)
	}
	if v := updateUser.DisplayName; v != nil {
		set, args = append(set, "display_name = ?"), append(args, *v)
	}
	if v := updateUser.Email; v != nil {
		set, args = append(set, "email = ?"), append(args, *v)
	}
	if v := updateUser.AvatarURL; v != nil {
		set, args = append(set, "avatar_url = ?"), append(args, *v)
	}
	args = append(args, updateUser.ID)

	query := `
		UPDATE users
		SET ` + joinStrings(set, ", ") + `
		WHERE id = ?
		RETURNING id,	created_at, updated_at, username, display_name, email, avatar_url
	`
	var user api.User
	if err := s.db.QueryRowContext(
		ctx, query, args...,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Username,
		&user.DisplayName,
		&user.Email,
		&user.AvatarURL,
	); err != nil {
		return nil, FormatError(err)
	}
	s.userCache.Set(&user)
	return &user, nil
}

func (s *Store) FindUser(ctx context.Context, findUser *api.FindUser) (*api.User, error) {
	if findUser.ID != nil {
		if user, ok := s.userCache.Get(); ok {
			return user.(*api.User), nil
		}
	}

	list, err := s.ListUsers(ctx, findUser)
	if err != nil {
		return nil, FormatError(err)
	}

	if len(list) == 0 {
		return nil, FormatError(err)
	}

	user := list[0]
	s.userCache.Set(user)
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
		return nil, FormatError(err)
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
			return nil, FormatError(err)
		}
		list = append(list, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, FormatError(err)
	}

	return list, nil
}

func (s *Store) UpsertUserSetting(ctx context.Context, upsert *api.UserSetting) (*api.UserSetting, error) {
	query := `
		INSERT INTO user_settings
		VALUES (?, ?, ?)
		ON CONFLICT(user_id, setting_key) DO UPDATE
		SET setting_value = EXCLUDED.setting_value
	`

	if _, err := s.db.ExecContext(ctx, query, &upsert.UserId, &upsert.SettingKey, &upsert.SettingValue); err != nil {
		return nil, FormatError(err)
	}

	return upsert, nil
}

func (s *Store) ListUserSettings(ctx context.Context, userId int) ([]*api.UserSetting, error) {
	query := `
		SELECT user_id, setting_key, setting_value
		FROM user_settings
		WHERE user_id = ?
	`

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, FormatError(err)
	}
	defer rows.Close()

	list := make([]*api.UserSetting, 0)
	for rows.Next() {
		var userSetting api.UserSetting
		if err := rows.Scan(
			&userSetting.UserId,
			&userSetting.SettingKey,
			&userSetting.SettingValue,
		); err != nil {
			return nil, FormatError(err)
		}
		list = append(list, &userSetting)
	}

	if err := rows.Err(); err != nil {
		return nil, FormatError(err)
	}

	return list, nil
}
