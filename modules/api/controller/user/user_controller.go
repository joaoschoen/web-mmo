package user

import (
	"net/http"
	"time"
	"web-mmo/modules/api/model"
	"web-mmo/modules/utils/db"
	"web-mmo/modules/utils/security"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// DUMMY DATA
var DummyUser = model.UserData{
	Username: "someID",
	Email:    "jon@doe.com",
	Password: "password1",
}

func Register(ctx echo.Context) error {
	println("starting")
	// BODY
	var user model.UserData
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Malformed data.")
	}
	// Empty data
	if user.Email == "" || user.Password == "" || user.Username == "" {
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	// Database connection
	conn, err := db.GetConnection(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Service currently unavailable")
	}
	defer conn.Close()

	// Duplicate check
	// Username
	rows, err := conn.QueryContext(ctx.Request().Context(), "SELECT username FROM \"user\" WHERE username = $1;", user.Username)
	if err != nil {
		log.Errorf(err.Error())
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	defer rows.Close()
	if rows.Next() == true {
		return ctx.JSON(http.StatusBadRequest, "The selected username is already in use.")
	}

	// Email
	rows, err = conn.QueryContext(ctx.Request().Context(), "SELECT email FROM \"user\" WHERE email = $1;", user.Email)
	defer rows.Close()
	if err != nil {
		log.Errorf(err.Error())
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	if rows.Next() == true {
		return ctx.JSON(http.StatusBadRequest, "The provided email is already in use.")
	}

	// Password minimum requirements
	if security.CheckPassword(user.Password) == false {
		return ctx.JSON(http.StatusBadRequest, "The provided password does not satisfy security standards.")
	}

	// Hashing
	hashedPassword, err := security.Encrypt(user.Password)
	if err != nil {
		println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	// Create User
	_, err = conn.ExecContext(ctx.Request().Context(), `
		INSERT INTO "user" (username,email,password,created_at)
		VALUES ($1,$2,$3,$4);`, user.Username, user.Email, string(hashedPassword), time.Now())
	if err != nil {
		println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func Login(ctx echo.Context) error {
	// BODY
	var user model.LoginCredentials
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Malformed data")
	}
	// Empty data
	if user.Email == "" || user.Password == "" {
		return ctx.JSON(http.StatusBadRequest, "Bad credentials")
	}

	// DATABASE REQUEST GOES HERE

	// CHECK DUPLICATE
	if user.Email == "alreadyIn@use.com" {
		return ctx.JSON(http.StatusUnprocessableEntity, "Email already in use")
	}
	// TODO
	// return token

	return ctx.JSON(http.StatusOK, nil)
}

// func PutUser(ctx echo.Context) error {
// 	// PARAM
// 	var id string
// 	id = ctx.Param("id")

// 	// BODY
// 	var user model.UnsafeUser
// 	err := ctx.Bind(&user)
// 	if err != nil {
// 		return ctx.String(http.StatusBadRequest, "Error while parsing received data")
// 	}
// 	// Empty data
// 	if user.Email == "" || user.Password == "" {
// 		return ctx.String(http.StatusBadRequest, "Error while parsing received data")
// 	}

// 	// DATABASE REQUEST GOES HERE

// 	// SIMULATED NOT FOUND
// 	if id == "404" {
// 		return ctx.JSON(http.StatusNotFound, "User not found")
// 	}

// 	// SIMULATED DUPLICATE
// 	if user.Email == "alreadyIn@use.com" {
// 		return ctx.JSON(http.StatusUnprocessableEntity, "Email already in use")
// 	}

// 	// BUILD RESPONSE
// 	response := model.PutUserResponse{
// 		Data: model.SafeUser{
// 			ID:    id,
// 			Email: user.Email,
// 		},
// 	}

// 	return ctx.JSON(http.StatusOK, response)
// }

// func DeleteUser(ctx echo.Context) error {
// 	// PARAM
// 	var id string
// 	id = ctx.Param("id")

// 	// SIMULATED NOT FOUND
// 	if id == "404" {
// 		return ctx.JSON(http.StatusNotFound, "User doesn't exist")
// 	}

// 	// DATABASE REQUEST GOES HERE

// 	//BUILD RESPONSE
// 	response := model.DeleteUserResponse{
// 		ID: id,
// 	}

// 	return ctx.JSON(http.StatusOK, response)
// }
