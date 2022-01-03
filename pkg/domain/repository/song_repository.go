package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/rrih/managedby/pkg/domain/entity"
	"github.com/rrih/managedby/pkg/infrastructure"
)

// FindAllSongs 全ての曲を返す
func FindAllSongs(db *sql.DB) ([]entity.Song, error) {
	rows, err := db.Query(
		`
			select
				id, registered_user_id, category_id, name, singer_name, composer_name,
				source, url, is_anime_video_dam, is_anime_video_joy, is_official_video_dam,
				is_official_video_joy, start_singing, deleted, created, modified
			from
				songs
			where
				deleted is null
		`,
	)
	if err != nil {
		return nil, err
	}
	song := entity.Song{}
	songs := []entity.Song{}
	for rows.Next() {
		var id, registeredUserID, categoryID, name, singerName, composerName, source string
		var url, startSinging, deleted *string
		var isAnimeVideoDam, isAnimeVideoJoy, isOfficialVideoDam, isOfficialVideoJoy bool
		var created, modified string
		err := rows.Scan(
			&id, &registeredUserID, &categoryID, &name, &singerName,
			&composerName, &source, &url, &isAnimeVideoDam, &isAnimeVideoJoy,
			&isOfficialVideoDam, &isOfficialVideoJoy, &startSinging, &deleted, &created, &modified,
		)
		if err != nil {
			panic(err)
		}
		song.ID = id
		song.RegisteredUserID = registeredUserID
		song.CategoryID = categoryID
		song.Name = name
		song.SingerName = singerName
		song.ComposerName = composerName
		song.Source = source
		song.URL = url
		song.IsAnimeVideoDam = isAnimeVideoDam
		song.IsAnimeVideoJoy = isAnimeVideoJoy
		song.IsOfficialVideoDam = isOfficialVideoDam
		song.IsOfficialVideoJoy = isOfficialVideoJoy
		song.StartSinging = startSinging
		song.Deleted = deleted
		song.Created = created
		song.Modified = modified
		songs = append(songs, song)
	}
	defer db.Close()
	return songs, nil
}

// FindSongByID songs.idから曲を取得する
func FindSongByID(songID int, db *sql.DB) (entity.Song, error) {
	row, err := db.Query(
		`
			select
				id, registered_user_id, category_id, name, singer_name, composer_name,
				source, url, is_anime_video_dam, is_anime_video_joy, is_official_video_dam,
				is_official_video_joy, start_singing, deleted, created, modified
			from
				songs
			where
				deleted is null
			and id = ?
		`, songID,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	var s entity.Song
	for row.Next() {
		err := row.Scan(
			&s.ID, &s.RegisteredUserID, &s.CategoryID, &s.Name, &s.SingerName, &s.ComposerName,
			&s.Source, &s.URL, &s.IsAnimeVideoDam, &s.IsAnimeVideoJoy, &s.IsOfficialVideoDam, &s.IsOfficialVideoJoy,
			&s.StartSinging, &s.Deleted, &s.Created, &s.Modified,
		)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	defer row.Close()
	return s, nil
}

// SaveSong 曲の保存
// TODO: ref: https://qiita.com/mizumizue/items/12d504eead84214af0e1
// こちらを参考にinsertしたレコードのidを取得して返す。
func SaveSong(s entity.Song, db *sql.DB) error {
	// TODO: 日本時間にする
	created, modified := time.Now().Format("2001-01-01 00:00:00"), time.Now().Format("2001-01-01 00:00:00")
	_, err := db.Exec(
		`
			insert into songs (
				registered_user_id, category_id, name, singer_name, composer_name,
				source, url, is_anime_video_dam, is_anime_video_joy, is_official_video_dam,
				is_official_video_joy, start_singing, deleted, created, modified
			) values (
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
			)
		`, s.RegisteredUserID, s.CategoryID, s.Name, s.SingerName, s.ComposerName,
		s.Source, s.URL, s.IsAnimeVideoDam, s.IsAnimeVideoJoy, s.IsOfficialVideoDam,
		s.IsOfficialVideoJoy, s.StartSinging, s.Deleted, created, modified,
	)
	return err
}

// UpdateSong 曲の更新
func UpdateSong(s entity.Song) {
	db := infrastructure.DbConn()
	// TODO: 日本時間にする
	modified := time.Now()
	_, err := db.Exec(
		`
			update
				songs
			set
			registered_user_id = ?, category_id = ?, name = ?, singer_name = ?, composer_name = ?,
			source = ?, url = ?, is_anime_video_dam = ?, is_anime_video_joy = ?, is_official_video_dam = ?,
			is_official_video_joy = ?, start_singing = ?, deleted = ?, modified = ?
			where
				id = ?
		`,
		s.RegisteredUserID, s.CategoryID, s.Name, s.SingerName, s.ComposerName,
		s.Source, s.URL, s.IsAnimeVideoDam, s.IsAnimeVideoJoy, s.IsOfficialVideoDam,
		s.IsOfficialVideoJoy, s.StartSinging, s.Deleted, modified,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// DeleteSong 曲削除
func DeleteSong(s entity.Song) {
	db := infrastructure.DbConn()
	modified := time.Now()
	deleted := time.Now()
	_, err := db.Exec(
		`
			update
				songs
			set
				deleted = ?, modified = ?
			where
				id = ?
		`,
		deleted, modified, s.ID,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
