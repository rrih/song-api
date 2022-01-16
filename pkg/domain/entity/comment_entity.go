package entity

// Comment エンティティ
type Comment struct {
	ID       string  `json:"id"`
	Text     string  `json:"text"`
	Deleted  *string `json:"deleted"`
	Created  string  `json:"created"`
	Modified string  `json:"modified"`
}
