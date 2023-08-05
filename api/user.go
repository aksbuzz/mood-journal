package api

import (
	"fmt"
	"net/mail"
	"time"
)

type User struct {
	ID           int    `json:"id"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	Username     string `json:"username"`
	DisplayName  string `json:"display_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	AvatarURL    string `json:"avatar_url"`
}

type FindUser struct {
	ID          *int
	UserName    *string
	DisplayName *string
	Email       *string
}
type UpdateUser struct {
	ID          *int
	Username    *string
	UpdatedAt   *int64
	DisplayName *string
	Email       *string
	AvatarURL   *string
}
type PatchUserRequest struct {
	Username    *string `json:"username"`
	UpdatedAt   *int64  `json:"updated_at"`
	DisplayName *string `json:"display_name"`
	Email       *string `json:"email"`
	AvatarURL   *string `json:"avatar_url"`
}

func (p *PatchUserRequest) Validate() error {
	if p.Username != nil && len(*p.Username) < 3 {
		return fmt.Errorf("username is too short")
	}
	if p.Username != nil && len(*p.Username) > 32 {
		return fmt.Errorf("username is too long")
	}
	if p.DisplayName != nil && len(*p.DisplayName) < 3 {
		return fmt.Errorf("display_name is too short")
	}
	if p.DisplayName != nil && len(*p.DisplayName) > 32 {
		return fmt.Errorf("display_name is too long")
	}
	if p.Email != nil {
		if _, err := mail.ParseAddress(*p.Email); err != nil {
			return fmt.Errorf("invalid email format")
		}
	}
	return nil
}

func (p *PatchUserRequest) GetUpdateUser(userId *int) *UpdateUser {
	currTime := time.Now().Unix()
	updateUser := &UpdateUser{
		ID:        userId,
		UpdatedAt: &currTime,
	}
	if p.Username != nil {
		updateUser.Username = p.Username
	}
	if p.DisplayName != nil {
		updateUser.DisplayName = p.DisplayName
	}
	if p.Email != nil {
		updateUser.Email = p.Email
	}
	if p.AvatarURL != nil {
		updateUser.AvatarURL = p.AvatarURL
	}
	return updateUser
}
