package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"golang-gocb-couchbase/configuration"
	_ "golang-gocb-couchbase/docs"
	"golang-gocb-couchbase/internal/folksdev-fiber-rest-api/application/controller"
	"golang-gocb-couchbase/internal/folksdev-fiber-rest-api/application/handler/user"
	"golang-gocb-couchbase/internal/folksdev-fiber-rest-api/application/query"
	"golang-gocb-couchbase/internal/folksdev-fiber-rest-api/application/repository"
	"golang-gocb-couchbase/internal/folksdev-fiber-rest-api/application/web"
	"golang-gocb-couchbase/internal/folksdev-fiber-rest-api/pkg/couchbase"
	"golang-gocb-couchbase/internal/folksdev-fiber-rest-api/pkg/server"
	"golang-gocb-couchbase/internal/folksdev-fiber-rest-api/pkg/utils"
	"golang-gocb-couchbase/internal/folksdev-fiber-rest-api/pkg/validation"
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

	byteOperationUtil := utils.NewByteOperationUtil()

	// couchbase connection
	couchbaseCluster := couchbase.ConnectCluster(configuration.CouchbaseHost,
		configuration.CouchbaseUsername,
		configuration.CouchbasePassword,
		byteOperationUtil.MbToBytes(uint(configuration.CouchbaseConnBufferSizeInMb)))

	// custom validator initializing
	customValidator := validation.NewCustomValidator(validator.New())
	customValidator.RegisterCustomValidation()

	// dependency injection
	userRepository := repository.NewUserRepository(couchbaseCluster, utils.NewExpirationUtil())
	userQueryService := query.NewUserQueryService(userRepository)
	userCommandHandler := user.NewCommandHandler(userRepository)
	userController := controller.NewUserController(userQueryService, userCommandHandler, customValidator)

	// Router initializing
	web.InitRouter(app, userController)

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
