package api

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
