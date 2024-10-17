// create_post.go

package crud

import (
	"KCLHack-PU-Back/database"
)

func CreatePostDB(post database.Post) (database.Post, error) {

	if err := database.DB.Create(&post).Error; err != nil {
		return database.Post{}, err
	}

	return post, nil
}
