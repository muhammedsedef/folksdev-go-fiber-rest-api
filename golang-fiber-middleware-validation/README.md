# Fiber & Middleware & Validation

This repository contains the source code discussed in the [Fiber & Middleware & Validation](https://youtube.com/playlist?list=PLCp1YoRkzkpY8YbNC3g503S1Q3Dsg7eaD&si=H5ll7GgITXfMYCV3) 
where we explored how to work with GoLang Fiber Framework, defining middleware, and performing validation on structs using the Validator library.

## Contents
- **Fiber Example**:
    - Inside the [main.go](main.go) directory, you can find sample code demonstrating how to set up a http server in our 
locale using the GoLang Fiber Framework.
- **Middleware Example**:
    - Inside the [main.go](main.go) directory, you can find sample code demonstrating how to define middleware using the
GoLang Fiber Framework.
- **Validation Example**:
    - Inside the [main.go](main.go) directory, you can find sample code demonstrating how to perform validation and handle 
validation error on structs using the Validator library.

## How to Use
- Before start to test example, be sure you're done [Requirements](../README.md) step 

1. **Middleware Example**:
    - To run the middleware example Run the command `go run main.go` to execute the example.
    - Visit `http://localhost:3000` in your browser or use postman collection to trigger endpoint 
    - Check application log on terminal
2. **Validation Example**:
    - To run the validation example Run the command `go run main.go` to execute the example.
    - Visit `http://localhost:3000` in your browser or use postman collection to trigger endpoint
    - If the struct model you send in the request is suitable for validations, you will receive a 200 success response message.
    - If the struct model you send in the request is not suitable with validations, you will receive a response message containing error details.

