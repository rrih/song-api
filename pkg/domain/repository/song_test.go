package repository

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rrih/managedby/pkg/domain/entity"
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

func Test_Success_FindSongByID(t *testing.T) {
	t.Run(
		"FindSongByIDが成功するケース",
		func(t *testing.T) {
			id := 1
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectQuery(
				regexp.QuoteMeta(
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
					`,
				),
			).
				WithArgs(id).
				WillReturnRows(sqlmock.NewRows([]string{
					"id", "registeredUserID", "categoryID", "name", "singerName", "composerName",
					"source", "url", "startSinging", "deleted", "isAnimeVideoDam", "isAnimeVideoJoy",
					"isOfficialVideoDam", "isOfficialVideoJoy", "created", "modified",
				}).AddRow(
					1, 1, 1, "三原色", "YOASOBI", "YOASOBI", "ahamoのCM", "https://www.youtube.com/watch?v=nhOhFOoURnE", 0, 0, 0, 0, "どこかで途切れた物語", nil, "2021-07-19 00:00:00", "2021-07-19 00:00:00",
				))

			_, err = FindSongByID(id, db)

			// assert
			if err != nil {
				t.Error(err.Error())
			}
		},
	)
}

func Test_Success_SaveSong(t *testing.T) {
	t.Run(
		"SaveSongが成功する",
		func(t *testing.T) {
			url := "https://example.com/xxxx/yyyy/zzzz"
			startSingingStr := "夜の運命埋め尽くす輝く夢となる"
			time := time.Now().Format("2001-01-01 00:00:00")
			s := entity.Song{
				// ID:                 1,
				RegisteredUserID:   "1",
				CategoryID:         "1",
				Name:               "FooBar",
				SingerName:         "Hoge Fuga",
				ComposerName:       "Piyo Piyo",
				Source:             "テレビのCM",
				URL:                &url,
				IsAnimeVideoDam:    true,
				IsAnimeVideoJoy:    true,
				IsOfficialVideoDam: true,
				IsOfficialVideoJoy: true,
				StartSinging:       &startSingingStr,
				Deleted:            nil,
				Created:            time,
				Modified:           time,
			}
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			// 期待
			mock.ExpectExec(regexp.QuoteMeta(
				`
					insert into songs (
						registered_user_id, category_id, name, singer_name, composer_name,
						source, url, is_anime_video_dam, is_anime_video_joy, is_official_video_dam,
						is_official_video_joy, start_singing, deleted, created, modified
					) values (
						?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
					)
				`,
			)).WithArgs(
				s.RegisteredUserID, s.CategoryID, s.Name, s.SingerName, s.ComposerName, s.Source,
				s.URL, s.IsAnimeVideoDam, s.IsAnimeVideoJoy, s.IsOfficialVideoDam, s.IsOfficialVideoJoy,
				s.StartSinging, s.Deleted, time, time,
			).WillReturnResult(sqlmock.NewResult(1, 1))

			// 実行
			err = SaveSong(s, db)

			// 結果
			if err != nil {
				t.Error(err.Error())
			}
		},
	)
}

func Test_Success_UpdateSong(t *testing.T) {}

func Test_Success_DeleteSong(t *testing.T) {}
