package model

import (
	"hello/database"
)

type UserModel struct {
	Id       int    `form:"id"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (user UserModel) TableName() string {
	return "user"
}

func (user UserModel) Insert() int {
	create := database.Db.Create(&user)
	if create.Error != nil {
		return 0
	}
	return 1
}

func (user UserModel) FindAll() []UserModel {
	var users []UserModel
	database.Db.Find(&users)
	return users
}

func (user UserModel) FindById() UserModel {
	database.Db.First(&user, user.Id)
	return user

}

func (user UserModel) DeleteOne() {
	database.Db.Delete(user)
}
