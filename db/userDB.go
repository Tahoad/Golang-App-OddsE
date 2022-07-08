package db

import "gorm.io/gorm"

type UserDB struct {
	gorm.Model
	//Id not use -> gorm will generate id automaticaly
	Code  string
	Name  string
	Email string
}

func Read(code string) (*UserDB, error) {
	db := DB()
	var userDB UserDB
	if err := db.First(&userDB, "Code = ?", code).Error; err != nil {
		return nil, err
	}

	return &userDB, nil
}

func Create(user UserDB) error {
	db := DB()
	userData := UserDB{
		Code:  user.Code,
		Name:  user.Name,
		Email: user.Email,
	}

	if err := db.Create(&userData).Error; err != nil {
		return err
	}

	return nil
}
