// create_post.go

package crud

import (
	"KCL-Hack2024-PU-Back/database"
)

func CreatePostDB(post database.Post) (database.Post, error) {

	if err := database.DB.Create(&post).Error; err != nil {
		return database.Post{}, err
	}

	return post, nil
}
