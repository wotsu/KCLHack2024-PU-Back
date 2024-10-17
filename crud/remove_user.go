package crud

import (
	"KCLHack-PU-Back/database"
)

func RemoveUser(userID uint) database.User {

	user := database.User{}
	database.DB.Where("id = ?", userID).Delete(&user)
	return user

}