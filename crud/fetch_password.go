package crud

import (
	"KCLHack-PU-Back/database"
)

func FetchPassword(userID uint) (string, error) {

	var user database.User

	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return user.Password, err
	}
	return user.Password, nil

}