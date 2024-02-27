package userM

import (
	"arooClasses/pkg/comman/userAuthModels"
	"errors"
	"fmt"
)

func Login(login userAuthModels.Login) (string, error) {
	// Extract credentials from the login struct
	email := login.Email
	userName := login.UserName
	password := login.Password

	emailOk, _ := validMailAddress(email)
	userNameOk := validateUserName(userName)
	passwordOk := validatePassword(password)

	if (emailOk || userNameOk) && passwordOk {
		// TODO: Store data into the database (replace with your implementation)
		// database.StoreUser(login)
		return "success", nil
	}

	return "", errors.New("Invalid credentials")
}
