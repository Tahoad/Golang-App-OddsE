package db

import "gorm.io/gorm"

type EmployeeDB struct {
	gorm.Model
	//Id not use -> gorm will generate id automaticaly
	Id       string
	Username string
	Password string
	Email    string
	Name     string
	Surname  string
}

func ReadEmp(code string) (*EmployeeDB, error) {
	db := DB()
	var employeeDB EmployeeDB
	if err := db.First(&employeeDB, "Code = ?", code).Error; err != nil {
		return nil, err
	}

	return &employeeDB, nil
}
func CreateEmp(emp EmployeeDB) error {
	db := DB()
	employeeData := EmployeeDB{
		Username: emp.Username,
		Password: emp.Password,
		Email:    emp.Email,
		Name:     emp.Name,
		Surname:  emp.Surname,
	}

	if err := db.Create(&employeeData).Error; err != nil {
		return err
	}

	return nil
}
