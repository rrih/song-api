package repository

import (
	"log"
	"time"

	"github.com/rrih/managedby/pkg/domain/entity"
	"github.com/rrih/managedby/pkg/infrastructure"
)

// FindAllCategories 全てのカテゴリーを返す
func FindAllCategories() []entity.Category {
	db := infrastructure.DbConn()
	rows, err := db.Query(
		`
			select
				id, parent_id, name, deleted, created, modified
			from
				songs
			where
				deleted is null
		`,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	category := entity.Category{}
	categories := []entity.Category{}
	for rows.Next() {
		var id, name string
		var parentID, deleted *string
		var created, modified string
		err := rows.Scan(
			&id, &parentID, &name, &deleted, &created, &modified,
		)
		if err != nil {
			panic(err)
		}
		category.ID = id
		category.ParentID = parentID
		category.Name = name
		category.Deleted = deleted
		category.Created = created
		category.Modified = modified
		categories = append(categories, category)
	}
	defer db.Close()
	return categories
}

// FindCategoryByID categories.idからカテゴリを取得する
func FindCategoryByID(id int) (entity.Category, error) {
	db := infrastructure.DbConn()
	row, err := db.Query(
		`
			select
				id, parent_id, name, deleted, created, modified
			from
				categories
			where
				deleted is null
			and id = ?
		`, id,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	var c entity.Category
	for row.Next() {
		err := row.Scan(
			&c.ID, &c.ParentID, &c.Name, &c.Deleted, &c.Created, &c.Modified,
		)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	defer row.Close()
	return c, nil
}

// SaveCategory カテゴリーの保存
func SaveCategory(c entity.Category) error {
	db := infrastructure.DbConn()
	// TODO: 日本時間にする
	created, modified := time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec(
		`
			insert into categories (
				parent_id, name, deleted, created, modified
			) values (
				?, ?, ?, ?, ?
			)
		`, c.ParentID, c.Name, c.Deleted, created, modified,
	)
	return err
}

// UpdateCategory カテゴリーの更新
// 不要な気がする
func UpdateCategory(c entity.Category) {
	db := infrastructure.DbConn()
	// TODO: 日本時間にする
	modified := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec(
		`
			update
				categories
			set
			parent_id = ?, name = ?, deleted = ?, modified = ?
			where
				id = ?
		`,
		c.ParentID, c.Name, c.Deleted, modified,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// DeleteCategory カテゴリー削除
// 不要？admin権限のみでいいかも
func DeleteCategory(c entity.Category) {
	db := infrastructure.DbConn()
	modified := time.Now().Format("2006-01-02 15:04:05")
	deleted := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec(
		`
			update
				categories
			set
				deleted = ?, modified = ?
			where
				id = ?
		`,
		deleted, modified, c.ID,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
