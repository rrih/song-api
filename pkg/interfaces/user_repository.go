package database

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query(
		"select id, name, email, password, created, modified from users order by id asc"
	)
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}
	// for rows.Next() {
		// var u domain.Users
		// err := 
	// }
}