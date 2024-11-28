package dao

import (
	"database/sql"
	"db/model"
)

type UserDAO struct {
	DB *sql.DB
}

func (dao *UserDAO) FindUserByName(name string) ([]model.User, error) {
	rows, err := dao.DB.Query("SELECT id, name, age FROM user WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (dao *UserDAO) InsertUser(user model.User) error {
	_, err := dao.DB.Exec("INSERT INTO user (id, name, age) VALUES (?, ?, ?)", user.ID, user.Name, user.Age)
	return err
}
