package models

import (
	"../config"
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	ID    int
	Email string `db:"email"`
}

// SaveUser - saves a user record to the db
func (u *User) SaveUser() (found User, err error) {
	user := User{Email: u.Email}

	record := config.DB.Save(&user)

	return user, record.Error
}
