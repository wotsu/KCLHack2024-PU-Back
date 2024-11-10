package crud

import (
	"KCLHack-PU-Back/database"
)

func FetchUserID(username string) (uint, error) {

	var user database.User

	if err := database.DB.Where("name = ?", username).First(&user).Error; err != nil {
		return user.ID, err
	}
	return user.ID, nil

}