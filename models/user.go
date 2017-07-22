package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	ID        int64     `db:"id"`
	email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
