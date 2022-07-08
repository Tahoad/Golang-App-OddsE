package users

import (
	"fmt"
	db "golang-app/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetEmployee(c echo.Context) error {
	// id := c.Param("id")
	// user := User{
	// 	Id:    id,
	// 	Name:  "John Doe",
	// 	Email: "john@gmail.com",
	// }

	empDB, err := db.ReadEmp(c.Param("id"))
	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	emp := Employee{
		Id:       empDB.Id,
		Username: empDB.Username,
		Password: empDB.Password,
		Name:     empDB.Name,
		Surname:  empDB.Surname,
		Email:    empDB.Email,
	}
	return c.JSON(http.StatusOK, emp)
}

func RegisterEmployee(c echo.Context) error {
	emp := Employee{}
	fmt.Print("c = ", c)
	if err := c.Bind(&emp); err != nil { //err is method return handler
		return err
	}

	employeeDB := db.EmployeeDB{
		Id:       emp.Id,
		Username: emp.Username,
		Password: emp.Password,
		Name:     emp.Name,
		Surname:  emp.Surname,
		Email:    emp.Email,
	}

	if err := db.CreateEmp(employeeDB); err != nil {
		return c.String(http.StatusExpectationFailed, "Create Fail.")
	}

	return c.JSON(http.StatusCreated, "created")

}

// func Update(c echo.Context) error {

// 	return c.String(http.StatusOK, "OK")
// }

// func Delete(c echo.Context) error {

// 	return c.String(http.StatusOK, "OK")
// }
// func Register(c echo.Context) error {
// 	emp := Employee{}
// 	if err := c.Bind(&emp); err != nil {
// 		return err
// 	}

// 	return c.String(http.StatusOK, "OK")
// }
