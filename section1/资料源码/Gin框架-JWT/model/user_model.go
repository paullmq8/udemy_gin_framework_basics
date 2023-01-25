package model

import "hello/database"

type UserModel struct {
	Id       int    `form:"id"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (user UserModel) TableName() string {
	return "user"
}

func (user UserModel) QueryByEmail() UserModel {
	database.Db.First(&user, user.Email)
	return user
}
