package entity

// ScoreComment エンティティ
type ScoreComment struct {
	ID        string `json:"id"`
	ScoreID   string `json:"score_id"`
	CommentID string `json:"comment_id"`
}
