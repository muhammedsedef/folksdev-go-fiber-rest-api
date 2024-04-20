package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"golang-project-layout-swagger/configuration"
	_ "golang-project-layout-swagger/docs"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/application/controller"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/application/handler/user"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/application/query"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/application/repository"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/application/web"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/pkg/server"
)

// @title			Folksdev Fiber Rest Api
// @version		1.0
// @description	This is a sample swagger for folksdev rest api
// @contact.name	Folksdev
// @contact.email	folksdev@gmail.com
func main() {
	// fiber framework http server
	app := fiber.New()

	app.Use(recover.New())

	configureSwaggerUi(app)

	userRepository := repository.NewUserRepository()
	userQueryService := query.NewUserQueryService(userRepository)
	userCommandHandler := user.NewCommandHandler(userRepository)
	userController := controller.NewUserController(userQueryService, userCommandHandler)

	// Router initializing
	web.InitRouter(app, userController)

	//customValidator := validation.NewCustomValidator(validator.New())

	// Start server
	server.NewServer(app).StartHttpServer()
}

func configureSwaggerUi(app *fiber.App) {
	if configuration.Env != "prod" {
		// Swagger injection
		app.Get("/swagger/*", swagger.HandlerDefault)

		// Root path to SwaggerUI redirection
		app.Get("/", func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusMovedPermanently).Redirect("/swagger/index.html")
		})
	}
}
