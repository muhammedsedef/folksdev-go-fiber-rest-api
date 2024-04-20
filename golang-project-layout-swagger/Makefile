preinstall:
	@command -v swag || go install github.com/swaggo/swag/cmd/swag@latest

build:
	@go build -tags dynamic

start:
	@go run main.go -tags dynamic

generate-swagger:
	@swag init


