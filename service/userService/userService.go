package userService

import (
	"errors"
	"fiber-gorm/database"
	"fiber-gorm/models/userModel"
)

type UserService struct{}

func (userService UserService) FindUser(id int, user *userModel.Users) error {
	database.Database.Db.Find(&user, "id= ?", id)
	if user.ID == 0 {
		return errors.New("User not Found!!!")
	}
	return nil
}
