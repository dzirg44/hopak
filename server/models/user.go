package models

import (
	"github.com/jinzhu/gorm"
)

// User structure
type User struct {
	UserID       uint   `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(100);not null"`
	Email        string `gorm:"type:varchar(100);not null;unique_index"`
	Password     string `gorm:"not null"`
	ProfileImage string `sql:"size:255"`
	Posts        []Post `gorm:"ForeignKey:UserID"`
}

// IsUserExist checks if user already exists
func IsUserExist(db *gorm.DB, email string) bool {
	user := User{}
	// Find a user with passed email
	result := db.Where("email = ?", email).First(&user).RecordNotFound()

	return result
}

// UserByID assigns a user by id
func UserByID(db *gorm.DB, user *User) bool {
	result := db.Where("user_id = ?", user.UserID).First(&user).RecordNotFound()

	return result
}
