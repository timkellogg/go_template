package models

import (
	"log"

	"../config"
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	ID                int
	Email             string `db:"email"`
	EncryptedPassword string `db:"encrypted_password"`
}

// SaveUser - saves a user record to the db
func (u *User) SaveUser(password string) (found User, err error) {
	encryptedPassword, err := Encrypt(password)
	if err != nil {
		log.Println(err)
	}

	user := User{Email: u.Email, EncryptedPassword: encryptedPassword}

	record := config.DB.Create(&user)

	return user, record.Error
}
