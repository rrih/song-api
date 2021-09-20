package entity

// Category エンティティ
type Category struct {
	ID       string  `json:"id"`
	ParentID *string `json:"parent_id"`
	Name     string  `json:"name"`
	Deleted  *string `json:"deleted"`
	Created  string  `json:"created"`
	Modified string  `json:"modified"`
}
