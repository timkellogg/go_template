package models

import (
	"../config"
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Email string `db:"email"`
}

// SaveUser - saves a user record to the db
func (u *User) SaveUser(email string) (found User, err error) {
	user := User{Email: email}

	record := config.DB.Save(&user)

	return user, record.Error
}
