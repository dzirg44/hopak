package models

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

// Post structure
type Post struct {
	PostID    uint   `gorm:"primary_key" json:"post_id"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_time"`
	Title     string `sql:"size:255;not null" json:"title"`
	Content   string `sql:"size:255;not null" json:"content"`
	Image     string `sql:"size:255;not null" json:"image_url"`
}

// Validate validates that all params were passed
func (post *Post) Validate() error {
	if post.UserID <= 0 {
		return errors.New("post user_id is not recognized")
	}

	if post.Title == "" {
		return errors.New("post title is not recognized")
	}

	if post.Content == "" {
		return errors.New("post content is not recognized")
	}

	if post.Image == "" {
		return errors.New("post image URL is not recognized")
	}

	return nil
}

// GetPostsByUserID assigns a list of posts to a specific user
func GetPostsByUserID(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Where("user_id = ?", userID).Find(&posts)

	if err.Error != nil {
		log.Fatal(err.Error)
		return nil, err.Error
	}

	return posts, nil
}

// CreateNewPost should create a new post
func CreateNewPost(db *gorm.DB, post *Post) error {

	validErr := post.Validate()
	if validErr != nil {
		return validErr
	}
	err := db.Create(&post)

	if err.Error != nil {
		log.Fatal(err.Error)
		return err.Error
	}

	return nil
}
