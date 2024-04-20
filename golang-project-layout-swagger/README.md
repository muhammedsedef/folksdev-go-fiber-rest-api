# Project Layout & Swagger

This repository contains the source code discussed in the [Project Layout & Swagger](https://www.youtube.com/watch?v=zTq6ei6E9uY)
Here, we talked about what the project layout structure should be like when developing a rest API with GoLang and how Swagger integration can be done with the API.

## Contents
- **Project Layout**:
    - Inside the [main.go](main.go) directory, you can find sample code demonstrating how to set up a http server in our
      locale using the GoLang Fiber Framework.
    - You can check project folders for seeing how to do dependency injection, routing, swagger implementation, starting http server and project layout
best practices
- **Swagger**:
    - Inside the [main.go](main.go) directory, you can find sample code demonstrating how to implement swagger using the
      gofiber swagger library.

## How to Use
- Before start to test example, be sure you're done [Requirements](../README.md) step

1. **Swagger**:
    - To show the swagger, run the command `go run main.go` to execute the example.
    - Visit `http://localhost:8080` in your browse
    - Check the localhost path on browser for redirecting http://localhost:8080/swagger/index.htm