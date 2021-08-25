package entity

type User struct {
	ID       int     `json: "id"`
	Name     string  `json: "name"`
	Email    string  `json: "email"`
	Password string  `json: "password"`
	IsAdmin  bool    `json: "is_admin"`
	Deleted  *string `json: "deleted"`
	Created  string  `json: "created"`
	Modified string  `json: "modified"`
}

type InsertedUser struct {
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json: "password"`
	IsAdmin  bool   `json: "is_admin"`
}

//
type LoginRequest struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

type LoginResponse struct {
	Token string `json: "token"`
}
