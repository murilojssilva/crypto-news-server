package main

import (
	"crypto-news-server/internal/data"
	"crypto-news-server/internal/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadPosts()
	data.LoadUsers()

	router := gin.Default()

	router.GET("/posts", handler.GetPosts)
	router.GET("/posts/:id", handler.GetPostsById)
	router.POST("/posts", handler.PostPosts)
	router.DELETE("/posts/:id", handler.DeletePostById)
	router.PUT("/posts/:id", handler.UpdatePostById)

	router.GET("/users", handler.GetUsers)
	router.GET("/users/:id", handler.GetUsersById)
	router.POST("/users", handler.PostUsers)
	router.DELETE("/users/:id", handler.DeleteUserById)
	router.PUT("/users/:id", handler.UpdateUserById)

	log.Println("Server running on http://localhost:3000")
	router.Run(":3000")
}
