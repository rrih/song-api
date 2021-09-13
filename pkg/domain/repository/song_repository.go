package repository

import (
	"log"

	"github.com/rrih/managedby/pkg/domain/entity"
	"github.com/rrih/managedby/pkg/infrastructure"
)

// FindAllSongs 全ての曲を返す
func FindAllSongs() []entity.Song {
	db := infrastructure.DbConn()
	rows, err := db.Query(
		"select id, registered_user_id, category_id, name, singer_name, composer_name, source, url, is_anime_video_dam, is_anime_video_joy, is_official_video_dam, is_official_video_joy, start_singing, deleted, created, modified from songs where deleted is null",
	)
	if err != nil {
		log.Fatal(err.Error())
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
	return songs
}
