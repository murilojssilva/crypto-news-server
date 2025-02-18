package handler

import (
	"crypto-news-server/internal/data"
	"crypto-news-server/internal/models"
	"crypto-news-server/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"posts": data.Posts,
	})
}

func PostPosts(c *gin.Context) {
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	if err := service.ValidatePostTitle(&newPost, data.Posts); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	newPost.ID = len(data.Posts) + 1
	newPost.Created_at = time.Now()
	newPost.Updated_at = time.Now()
	data.Posts = append(data.Posts, newPost)

	data.SavePost()

	c.JSON(http.StatusCreated, newPost)
}

func GetPostsById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for _, p := range data.Posts {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Post not found",
	})
}

func DeletePostById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}
	for i, p := range data.Posts {
		if p.ID == id {
			data.Posts = append(data.Posts[:i], data.Posts[i+1:]...)
			data.SavePost()
			c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func UpdatePostById(c *gin.Context) {
	idParam := c.Param("id")

    id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"erro": err.Error(),
		})
		return
	}

	var updatedPost models.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})

		return
	}

	var existingPost *models.Post
	for i := range data.Posts {
		if data.Posts[i].ID == id {
			existingPost = &data.Posts[i]
			break
		}
	}

	if existingPost == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	if updatedPost.Title != "" {
		if err := service.ValidatePostTitle(&updatedPost, data.Posts); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
			return
		}
	} else {
		updatedPost.Title = existingPost.Title
	}

	if updatedPost.Subtitle == "" {
		updatedPost.Subtitle = existingPost.Subtitle
	}
	if updatedPost.Content == "" {
		updatedPost.Content = existingPost.Content
	}
	updatedPost.Written_by = existingPost.Written_by
	updatedPost.Created_at = existingPost.Created_at

	for i, p := range data.Posts {
		if p.ID == id {
			data.Posts[i] = updatedPost
			data.Posts[i].ID = id
			data.Posts[i].Updated_at = time.Now()
			data.SavePost()
			c.JSON(http.StatusCreated, gin.H{"message": "post updated"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})
}