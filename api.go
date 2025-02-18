package handler

import (
	"crypto-news-server/internal/data"
	"crypto-news-server/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

// setupRouter inicializa as rotas do servidor
func setupRouter() *gin.Engine {
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

	return router
}

// Handler é a função necessária para a Vercel processar requisições
func Handler(w http.ResponseWriter, r *http.Request) {
	setupRouter().ServeHTTP(w, r)
}
