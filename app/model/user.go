package model

import (
	"golang.org/x/crypto/bcrypt"
)

// defining a model struct
type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null;unique"`
	Password  string `json:"password" gorm:"not null"`
	IsEnabled bool   `json:"is_enabled"`
}

func Register(user User) (User, error) {
	// encrypting the password with bcrypt
	encrypted, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	// overriding the values
	user.IsEnabled = true
	user.Password = string(encrypted)

	// creating the user
	tx := DB.Create(&user)
	return user, tx.Error
}

func GetUserByEmail(email string) (User, error) {
	// getting the user form the db with the specified email
	var user User
	tx := DB.Where("email = ?", email).First(&user)
	return user, tx.Error
}
