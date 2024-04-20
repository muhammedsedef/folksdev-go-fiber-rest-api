package web

import (
	"github.com/gofiber/fiber/v2"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/application/controller"
	"net/http"
)

func InitRouter(app *fiber.App, userController controller.IUserController) {

	app.Get("/healthcheck", func(context *fiber.Ctx) error {
		return context.SendStatus(http.StatusOK)
	})

	folksdevRouteGroup := app.Group("/api/v1/folksdev")

	folksdevRouteGroup.Get("/user", userController.GetUser)
	folksdevRouteGroup.Post("/user", userController.Save)
	folksdevRouteGroup.Get("/user/:userId", userController.GetUserById)
}
