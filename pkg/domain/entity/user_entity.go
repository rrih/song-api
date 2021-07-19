package entity

type User struct {
	ID       int     `json: "id"`
	Name     string  `json: "string"`
	Email    string  `json: "string"`
	Password string  `json: "string"`
	IsAdmin  bool    `json: "bool"`
	Deleted  *string `json: "string"`
	Created  string  `json: "string"`
	Modified string  `json: "string"`
}

type InsertedUser struct {
	Name     string `json: "string"`
	Email    string `json: "string"`
	Password string `json: "string"`
	IsAdmin  bool   `json: "bool"`
}
