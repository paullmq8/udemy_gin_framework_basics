package model

import (
	"hello/database"
	"log"
)

type UserModel struct {
	Id       int    `form:"id"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (user UserModel) Insert() int {
	result, e := database.Db.Exec("insert into `user` (id,email,password) values (?,?,?);", user.Id, user.Email, user.Password)
	if e != nil {
		log.Panicln("添加用户失败")
	}
	i, _ := result.LastInsertId()
	return int(i)
}

func (user UserModel) FindAll() []UserModel {
	rows, e := database.Db.Query("select * from `user`")
	if e != nil {
		log.Panicln("查询数据失败")
	}

	var users []UserModel

	for rows.Next() {
		var u UserModel
		if e := rows.Scan(&u.Id, &u.Email, &u.Password); e == nil {
			users = append(users, u)
		}
	}
	return users
}

func (user UserModel) FindById() UserModel {
	row := database.Db.QueryRow("select * from `user` where id=?", user.Id)

	if e := row.Scan(&user.Id, &user.Email, &user.Password); e != nil {
		log.Panicln("绑定发生错误", e.Error())
	}
	return user

}

func (user UserModel) DeleteOne() {
	if _, e := database.Db.Exec("delete from `user` where id=?", user.Id); e != nil {
		log.Panicln("删除错误", e.Error())
	}
}
