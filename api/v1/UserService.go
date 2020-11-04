package service

import "github.com/gin-gonic/gin"

type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

func (userService *UserService) Create(c *gin.Context) {

}
func (userService *UserService) Login(c *gin.Context) {

}
func (userService *UserService) Edit(c *gin.Context) {

}

func (userService *UserService) Delete(c *gin.Context) {

}
func (userService *UserService) PostArticles(c *gin.Context) {

}

func (userService *UserService) EditArticles(c *gin.Context) {

}
func (userService *UserService) DeleteArticles(c *gin.Context) {

}
