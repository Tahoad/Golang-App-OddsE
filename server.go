package main

import (
	u "golang-app/users"

	db "golang-app/db"

	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	db.DB()
	db.Migrate()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"response" : "success"}`)
	})
	e.GET("/users/:id", u.GetUser)

	e.POST("/users", u.Save)

	e.PUT("/users/:id", u.Update)

	e.DELETE("/users", u.Delete)

	e.Logger.Fatal(e.Start(":1234"))

}
