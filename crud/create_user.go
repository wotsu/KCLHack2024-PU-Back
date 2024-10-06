// create_user.go

package crud

import (
	"KCLHack-PU-Back/database"
)

func CreateUserDB(user database.User) database.User {

	database.DB.Create(&user)
	return user

}
