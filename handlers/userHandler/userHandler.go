package userHandler

import (
	"fiber-gorm/database"
	"fiber-gorm/models/userModel"
	"fiber-gorm/response/userResponse"
	"fiber-gorm/service/userService"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler struct{}

func (userHandler UserHandler) CreateUser(c *fiber.Ctx) error {
	var user userModel.Users
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	ResponseUser := userResponse.UserResponse{}.CreateUserResponse(user)
	return c.Status(http.StatusCreated).JSON(ResponseUser)
}

func (userHandler UserHandler) GetAllUsers(c *fiber.Ctx) error {
	var users []userModel.Users
	database.Database.Db.Find(&users)
	if len(users) == 0 {
		return c.Status(404).JSON("User tapylmady!!!")
	}
	ResponseUser := []userResponse.UserResponse{}
	for _, user := range users {
		response := userResponse.UserResponse{}.CreateUserResponse(user)
		ResponseUser = append(ResponseUser, response)
	}
	return c.Status(200).JSON(ResponseUser)
}

func (userHandler UserHandler) GetUser(c *fiber.Ctx) error {
	var user userModel.Users
	userId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).JSON("user not Found :id!=integer !!!")
	}
	err = userService.UserService{}.FindUser(userId, &user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	ResponseUser := userResponse.UserResponse{}.CreateUserResponse(user)
	return c.Status(200).JSON(ResponseUser)
}

func (userHandler UserHandler) UpdateUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).JSON("user not Found :id!=integer !!!")
	}
	var user userModel.Users
	err = userService.UserService{}.FindUser(userId, &user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var updateUser userResponse.UpdateUser
	if err = c.BodyParser(&updateUser); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	user.FirstName = updateUser.FirstName
	user.LastName = updateUser.LastName
	database.Database.Db.Save(&user)
	ResponseUser := userResponse.UserResponse{}.CreateUserResponse(user)
	return c.Status(200).JSON(ResponseUser)
}

func (userHandler UserHandler) DeleteUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	var user userModel.Users
	if err != nil {
		return c.Status(404).JSON("user not Found :id!=integer !!!")
	}
	err = userService.UserService{}.FindUser(userId, &user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Unscoped().Delete(&user)
	return c.Status(200).JSON("user Deleted Successfully")

}
