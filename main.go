package main

import (
	service "go-blog/api/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//版本号v1
	v1 := router.Group("/v1")
	{
		userService := service.NewUserService()
		v1.POST("/users", userService.Create)
		v1.GET("/tokens", userService.Login)
		v1.PUT("/users/:user_id", userService.Edit)
		v1.DELETE("/users/:user_id", userService.Delete)
		v1.POST("/users/:user_id/articles", userService.PostArticles)
		v1.PUT("/users/:user_id/articles/:article_id", userService.EditArticles)
		v1.DELETE("/users/:user_id/articles/:article_id", userService.DeleteArticles)
	}

	router.Run(":8080")
}
