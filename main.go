package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"net/http"
)

var validate = validator.New()

type ErrorResponse struct {
	Status      int32                 `json:"status"`
	ErrorDetail []ErrorResponseDetail `json:"errorDetail"`
}

type ErrorResponseDetail struct {
	FieldName   string `json:"fieldName"`
	Description string `json:"description"`
}

type UserCreteRequest struct {
	FirstName string `json:"firstName" validate:"required,min=2"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required,min=8,max=16"`
	Age       int32  `json:"age" validate:"required,acceptAge"`
}

type CustomValidationError struct {
	HasError bool
	Field    string
	Tag      string
	Param    string
	Value    interface{}
}

func Validate(data interface{}) []CustomValidationError {
	var customValidationError []CustomValidationError

	if errors := validate.Struct(data); errors != nil {
		for _, fieldError := range errors.(validator.ValidationErrors) {
			var cve CustomValidationError
			cve.HasError = true
			cve.Field = fieldError.Field()
			cve.Tag = fieldError.Tag()
			cve.Param = fieldError.Param()
			cve.Value = fieldError.Value()
			customValidationError = append(customValidationError, cve)
		}
	}

	return customValidationError
}

func main() {
	// fiber framework http server
	app := fiber.New()

	// error description mapping
	validationErrorDescriptionMap := map[string]string{
		"min":       "Your value should be greater than ",
		"required":  "Your value is mandatory",
		"acceptAge": "Your value should be greater than 18",
	}

	app.Use(recover.New())

	// middleware
	app.Use(func(ctx *fiber.Ctx) error {
		fmt.Printf("Hello client, you're call my %s%s AND Method: %s\n", ctx.BaseURL(), ctx.Request().RequestURI(), ctx.Request().Header.Method())
		return ctx.Next()
	})

	// middleware for custom endpoint
	app.Use("/user", func(ctx *fiber.Ctx) error {
		correlationId := ctx.Get("X-CorrelationId")

		if correlationId == "" {
			return ctx.Status(http.StatusBadRequest).JSON("You have to send correlationId")
		}

		_, err := uuid.Parse(correlationId)

		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON("CorrelationId is must be guid")
		}

		ctx.Locals("correlationId", correlationId)
		return ctx.Next()
	})

	//recover middleware example
	app.Use(func(ctx *fiber.Ctx) error {
		//panic("The app is not crashing!!! It still work thanks to recover")
		fmt.Println("hello recover middleware test")
		return ctx.Next()
	})

	//custom validation registering
	validate.RegisterValidation("acceptAge", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() >= 18
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		fmt.Println("hello first get endpoint")
		return ctx.SendString("hello my first get endpoint")
	})

	app.Get("/user/:userId", func(ctx *fiber.Ctx) error {
		userIdParam := ctx.Params("userId")

		fmt.Printf("This is your userId -> %s\n", userIdParam)
		return ctx.SendString(fmt.Sprintf("your userId -> %s", userIdParam))
	})

	app.Post("/user", func(ctx *fiber.Ctx) error {
		fmt.Printf("hello correlationId -> %s\n", ctx.Locals("correlationId"))

		var request UserCreteRequest
		err := ctx.BodyParser(&request)

		if err != nil {
			fmt.Printf("There was an error while binding json - ERROR: %v\n", err.Error())
			return err
		}

		if errors := Validate(request); len(errors) > 0 && errors[0].HasError {
			var errorResponse ErrorResponse
			var errorDetailList []ErrorResponseDetail

			for _, validationError := range errors {
				var errorDetail ErrorResponseDetail
				errorDetail.FieldName = validationError.Field
				errorDetail.Description = fmt.Sprintf("%s field has en error because %s%s", validationError.Field, validationErrorDescriptionMap[validationError.Tag], validationError.Param)
				errorDetailList = append(errorDetailList, errorDetail)
			}
			errorResponse.Status = http.StatusBadRequest
			errorResponse.ErrorDetail = errorDetailList

			return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
		}

		responseMessage := fmt.Sprintf("%s created successfully", request.FirstName)
		return ctx.Status(http.StatusOK).JSON(responseMessage)
	})

	if err := app.Listen(":3000"); err != nil {
		fmt.Printf("Cannot start server")
	}
}
