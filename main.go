package main

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// // User
// type User struct {
// 	Name  string `json:"name" xml:"name"`
// 	Email string `json:"email" xml:"email"`
// }

// // Handler
// func(c echo.Context) error {
//   u := &User{
//     Name:  "Jon",
//     Email: "jon@labstack.com",
//   }
//   return c.JSON(http.StatusOK, u)
// }

type (
	User struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	spew.Dump("Hello")
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/users", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
