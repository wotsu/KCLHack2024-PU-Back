package crud

import (
	"KCLHack-PU-Back/database"
)

func FetchUsers(users []database.User) []database.User {

	database.DB.Find(&users)
	return users

}