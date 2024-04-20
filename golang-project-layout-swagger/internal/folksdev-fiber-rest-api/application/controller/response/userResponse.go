package response

import "golang-project-layout-swagger/internal/folksdev-fiber-rest-api/domain"

type UserResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       int32  `json:"age"`
}

func ToUserResponse(user *domain.User) UserResponse {
	return UserResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}
}

func ToUserResponseList(users []*domain.User) []UserResponse {
	var response = make([]UserResponse, 0)

	for _, user := range users {
		response = append(response, ToUserResponse(user))
	}

	return response
}
