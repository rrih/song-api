package entity

// Song 曲エンティティ
type Song struct {
	ID                 string  `json:"id"`
	RegisteredUserID   string  `json:"registered_user_id"`
	CategoryID         string  `json:"category_id"`
	Name               string  `json:"name"`
	SingerName         string  `json:"singer_name"`
	ComposerName       string  `json:"composer_name"`
	Source             string  `json:"source"`
	URL                *string `json:"url"`
	IsAnimeVideoDam    bool    `json:"is_anime_video_dam"`
	IsAnimeVideoJoy    bool    `json:"is_anime_video_joy"`
	IsOfficialVideoDam bool    `json:"is_official_video_dam"`
	IsOfficialVideoJoy bool    `json:"is_official_video_joy"`
	StartSinging       *string `json:"start_singing"`
	Deleted            *string `json:"deleted"`
	Created            string  `json:"created"`
	Modified           string  `json:"modified"`
}
