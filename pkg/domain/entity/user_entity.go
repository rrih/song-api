package entity

// TODO: あとで /pkg/domain 配下に移動する
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
