package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_Success_PhysicalDeleteUser(t *testing.T) {
	t.Run(
		"PhysicalDeleteUserが成功するケース",
		func(t *testing.T) {
			id := 1
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectExec(regexp.QuoteMeta(
				`
					delete from
						users
					where
						id = ?
				`,
			)).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

			err = PhysicalDeleteUser(id, db)

			if err != nil {
				t.Error(err.Error())
			}
		},
	)
}
