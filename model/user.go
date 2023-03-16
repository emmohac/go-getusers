package model

import (
	"html"
	"strings"
	"user_api/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Admin    bool   `gorm:"default: false" json:"admin"`
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// gorm hook
func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func FindAll() ([]User, error) {
	var users []User
	err := database.Database.Find(&users).Error
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

func (user *User) Update() (*User, error) {
	err := database.Database.Find(&user).Error

	if err != nil {
		return &User{}, err
	}

	err = database.Database.Updates(&user).Error

	if err != nil {
		return &User{}, err
	}

	return user, nil
}
