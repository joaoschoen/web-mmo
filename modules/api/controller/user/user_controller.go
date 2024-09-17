package user

import (
	"net/http"
	"web-mmo/modules/api/model"

	"github.com/labstack/echo/v4"
)

// DUMMY DATA
var DummyUser = model.UserData{
	ID:       "someID",
	Email:    "jon@doe.com",
	Password: "password1",
}

func Register(context echo.Context) error {
	// BODY
	var user model.LoginCredentials
	err := context.Bind(&user)
	if err != nil {
		return context.JSON(http.StatusBadRequest, "Malformed data")
	}
	// Empty data
	if user.Email == "" || user.Password == "" {
		return context.JSON(http.StatusBadRequest, "Bad credentials")
	}

	// DATABASE REQUEST GOES HERE

	// CHECK DUPLICATE
	if user.Email == "alreadyIn@use.com" {
		return context.JSON(http.StatusUnprocessableEntity, "Email already in use")
	}

	return context.JSON(http.StatusCreated, nil)
}

func Login(context echo.Context) error {
	// BODY
	var user model.LoginCredentials
	err := context.Bind(&user)
	if err != nil {
		return context.JSON(http.StatusBadRequest, "Malformed data")
	}
	// Empty data
	if user.Email == "" || user.Password == "" {
		return context.JSON(http.StatusBadRequest, "Bad credentials")
	}

	// DATABASE REQUEST GOES HERE

	// CHECK DUPLICATE
	if user.Email == "alreadyIn@use.com" {
		return context.JSON(http.StatusUnprocessableEntity, "Email already in use")
	}
	// TODO
	// return token

	return context.JSON(http.StatusOK, nil)
}

// func PutUser(context echo.Context) error {
// 	// PARAM
// 	var id string
// 	id = context.Param("id")

// 	// BODY
// 	var user model.UnsafeUser
// 	err := context.Bind(&user)
// 	if err != nil {
// 		return context.String(http.StatusBadRequest, "Error while parsing received data")
// 	}
// 	// Empty data
// 	if user.Email == "" || user.Password == "" {
// 		return context.String(http.StatusBadRequest, "Error while parsing received data")
// 	}

// 	// DATABASE REQUEST GOES HERE

// 	// SIMULATED NOT FOUND
// 	if id == "404" {
// 		return context.JSON(http.StatusNotFound, "User not found")
// 	}

// 	// SIMULATED DUPLICATE
// 	if user.Email == "alreadyIn@use.com" {
// 		return context.JSON(http.StatusUnprocessableEntity, "Email already in use")
// 	}

// 	// BUILD RESPONSE
// 	response := model.PutUserResponse{
// 		Data: model.SafeUser{
// 			ID:    id,
// 			Email: user.Email,
// 		},
// 	}

// 	return context.JSON(http.StatusOK, response)
// }

// func DeleteUser(context echo.Context) error {
// 	// PARAM
// 	var id string
// 	id = context.Param("id")

// 	// SIMULATED NOT FOUND
// 	if id == "404" {
// 		return context.JSON(http.StatusNotFound, "User doesn't exist")
// 	}

// 	// DATABASE REQUEST GOES HERE

// 	//BUILD RESPONSE
// 	response := model.DeleteUserResponse{
// 		ID: id,
// 	}

// 	return context.JSON(http.StatusOK, response)
// }
