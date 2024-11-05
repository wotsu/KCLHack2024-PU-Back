package crud

import (
	"KCLHack-PU-Back/database"
)

func FetchPosts(posts []database.Post) []database.Post {

	database.DB.Find(&posts)
	return posts

}