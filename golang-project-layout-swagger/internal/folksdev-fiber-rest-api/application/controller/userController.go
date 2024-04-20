package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/application/controller/request"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/application/controller/response"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/application/handler/user"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/application/query"
	"net/http"
)

type IUserController interface {
	Save(ctx *fiber.Ctx) error
	GetUserById(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
}

type userController struct {
	userQueryService   query.IUserQueryService
	userCommandHandler user.ICommandHandler
}

func NewUserController(userQueryService query.IUserQueryService, userCommandHandler user.ICommandHandler) IUserController {
	return &userController{
		userQueryService:   userQueryService,
		userCommandHandler: userCommandHandler,
	}
}

// Save godoc

//	@Summary		This method used for saving new user
//	@Description	saving new user
//
// @Param requestBody body request.UserCreteRequest nil "Handle Request Body"
//
//	@Tags			User
//	@Accept			json
//	@Produce		json
//
// @Success 200
//
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/api/v1/folksdev/user [post]
func (c *userController) Save(ctx *fiber.Ctx) error {
	var req request.UserCreteRequest
	err := ctx.BodyParser(&req)

	if err != nil {
		fmt.Printf("userController.Save ERROR -> There was an error while binding json - ERROR: %v\n", err.Error())
		return err
	}

	fmt.Printf("userController.Save STARTED with request: %#v\n", req)

	if err = c.userCommandHandler.Save(ctx.UserContext(), req.ToCommand()); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON("User Created Successfully")
}

// GetUserById godoc
//
//	@Summary		This method get user by given id
//	@Description	get user by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string	true	"userId"
//
// @Success 200 {object} response.UserResponse
//
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/api/v1/folksdev/user/{userId} [get]
func (c *userController) GetUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	if userId == "" {
		return ctx.Status(http.StatusBadRequest).JSON("userId cannot be empty")
	}

	fmt.Printf("userController.GetUserById STARTED with userId: %s\n", userId)

	user, err := c.userQueryService.GetById(ctx.UserContext(), userId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ToUserResponse(user))
}

// GetUser godoc
//
//	@Summary		This method used for get all users
//	@Description	get all users
//	@Tags			User
//	@Accept			json
//	@Produce		json
//
// @Success 200 {object} []response.UserResponse
//
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/api/v1/folksdev/user [get]
func (c *userController) GetUser(ctx *fiber.Ctx) error {
	fmt.Printf("userController.GetUser INFO - Started \n")

	users, err := c.userQueryService.Get(ctx.UserContext())

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ToUserResponseList(users))
}
