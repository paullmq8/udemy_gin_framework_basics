package model

import (
	"hello/database"
	"log"
)

type UserModel struct {
	Id       int64  `form:"id"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (user *UserModel) Save() int64 {
	result, e := database.Db.Exec("insert into user (email,password) value (?,?);", user.Email, user.Password)

	if e != nil {
		log.Panicln(e.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln(err.Error())
	}
	return id
}

func (user *UserModel) QueryByEmail() UserModel {
	u := UserModel{}
	row := database.Db.QueryRow("select * from user where email = ?;", user.Email)
	e := row.Scan(&u.Id, &u.Email, &u.Password)
	if e != nil {
		log.Panicln(e)
	}
	return u
}
