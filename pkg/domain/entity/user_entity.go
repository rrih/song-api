package entity

// User ...
type User struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	IsAdmin  bool    `json:"is_admin"`
	Deleted  *string `json:"deleted"`
	Created  string  `json:"created"`
	Modified string  `json:"modified"`
}

// InsertedUser はユーザー登録用
type InsertedUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

// LoginRequest はログインリクエスト用
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse はログインレスポンス用
type LoginResponse struct {
	Token  string `json:"token"`
	UserID string `json:"login_user_id"`
}

type AuthObject struct {
	Token string `json:"token"`
}
