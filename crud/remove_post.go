package crud

import (
	"KCLHack-PU-Back/database"
)

func RemovePost(postID uint) database.Post {

	post := database.Post{}
	database.DB.Where("id = ?", postID).Delete(&post)
	return post

}