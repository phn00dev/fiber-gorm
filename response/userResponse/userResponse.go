package userResponse

import (
	"fiber-gorm/models/userModel"
)

type UserResponse struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (userResponse UserResponse) CreateUserResponse(user userModel.Users) UserResponse {
	return UserResponse{Id: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}
