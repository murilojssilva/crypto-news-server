package service

import (
	"crypto-news-server/internal/models"
	"errors"
)

func ValidatePostTitle(post *models.Post, existingPosts []models.Post) error {
	for _, t := range existingPosts {
		if t.Title == post.Title {
			return errors.New("já há uma notícia com esse título")
		}
	}
	return nil
}