package request

import "golang-project-layout-swagger/internal/folksdev-fiber-rest-api/application/handler/user"

type UserCreteRequest struct {
	FirstName string `json:"firstName" validate:"required,min=2"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required,min=8,max=16"`
	Age       int32  `json:"age" validate:"required,acceptAge"`
}

func (req *UserCreteRequest) ToCommand() user.Command {
	return user.Command{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Age:       req.Age,
	}
}
