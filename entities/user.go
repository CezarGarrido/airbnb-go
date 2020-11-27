package entities

import (
	"errors"
	"strings"
)

type User struct {
	Base
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Validate :
func (user *User) Validate() error {
	errorList := []string{}

	if len(strings.Trim(user.Email, " ")) <= 0 {
		errorList = append(errorList, "Email is required;")
	}

	if len(strings.Trim(user.FirstName, " ")) <= 0 {
		errorList = append(errorList, "FirstName is required;")
	}

	if len(strings.Trim(user.Password, " ")) <= 0 {
		errorList = append(errorList, "Password is required")
	}

	if len(errorList) > 0 {
		return errors.New(strings.Join(errorList, " "))
	}
	return nil
}
