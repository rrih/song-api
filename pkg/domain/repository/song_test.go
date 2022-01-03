package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_Success_FindAllSongs(t *testing.T) {
	// see: https://uzimihsr.github.io/post/2021-04-30-golang-test-with-go-sqlmock/

	t.Run(
		"FindAllSongsが成功するケース",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectQuery(regexp.QuoteMeta(`
				select
					id, registered_user_id, category_id, name, singer_name, composer_name,
					source, url, is_anime_video_dam, is_anime_video_joy, is_official_video_dam,
					is_official_video_joy, start_singing, deleted, created, modified
				from
					songs
				where
					deleted is null
			`)).WillReturnRows(sqlmock.NewRows([]string{
				"id", "registeredUserID", "categoryID", "name", "singerName", "composerName", "source", "url", "startSinging", "deleted", "isAnimeVideoDam", "isAnimeVideoJoy", "isOfficialVideoDam", "isOfficialVideoJoy", "created", "modified",
			}).AddRow(
				1, 1, 1, "三原色", "YOASOBI", "YOASOBI", "ahamoのCM", "https://www.youtube.com/watch?v=nhOhFOoURnE", 0, 0, 0, 0, "どこかで途切れた物語", nil, "2021-07-19 00:00:00", "2021-07-19 00:00:00",
			))

			_, err = FindAllSongs(db)

			// assert
			if err != nil {
				t.Error(err.Error())
			}
		},
	)
}

func Test_Success_FindSongByID(t *testing.T) {}

func Test_Success_SaveSong(t *testing.T) {}

func Test_Success_UpdateSong(t *testing.T) {}

func Test_Success_DeleteSong(t *testing.T) {}
