package main

import (
	"crypto-news-server/internal/data"
	"crypto-news-server/internal/handler"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadPosts()
	data.LoadUsers()

	router := gin.Default()

	// Configuração de CORS
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"}, // Permite essas origens
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Métodos permitidos
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Cabeçalhos permitidos
		AllowCredentials: true,
	}

	// Aplica o middleware de CORS com a configuração
	router.Use(cors.New(corsConfig))

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	log.Println("Server running on port:", port)
	router.Run(":" + port)
}
