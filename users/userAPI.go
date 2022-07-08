package users

import (
	db "golang-app/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

//this function is the same as text line 12
func GetUser(c echo.Context) error {
	// id := c.Param("id")
	// user := User{
	// 	Id:    id,
	// 	Name:  "John Doe",
	// 	Email: "john@gmail.com",
	// }

	userDB, err := db.Read(c.Param("id"))
	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	user := User{
		Id:    userDB.Code,
		Name:  userDB.Name,
		Email: userDB.Email,
	}
	return c.JSON(http.StatusOK, user)
}

func Save(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil { //err is method return handler
		return err
	}

	userDB := db.UserDB{
		Code:  user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	if err := db.Create(userDB); err != nil {
		return c.String(http.StatusExpectationFailed, "Create Fail.")
	}

	return c.JSON(http.StatusCreated, "created")

}

func Update(c echo.Context) error {

	return c.String(http.StatusOK, "OK")
}

func Delete(c echo.Context) error {

	return c.String(http.StatusOK, "OK")
}
