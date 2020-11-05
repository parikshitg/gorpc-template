package models

import "gorm.io/gorm"

// Db - gorm Db instance
var Db *gorm.DB

// User - user structure
type User struct {
	Name     string
	Phone    string
	Email    string
	Password string
}

// CreateUser - Creates a user in DB
func CreateUser(name, phone, email, password string) {
	Db.Create(&User{Name: name, Phone: phone, Email: email, Password: password})
}

// ExistingUser - checks if user exists in DB or not
func ExistingUser(email, password string) bool {
	var u User
	Db.Where("email = ? AND password = ?", email, password).First(&u)
	if email != u.Email && password != u.Password {
		return false
	}
	return true
}

// ListUser - returns a list of users
func ListUser(email, password string) []User {
	var u []User
	Db.Find(&u)
	return u
}
