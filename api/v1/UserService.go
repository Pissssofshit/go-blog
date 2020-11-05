package service

import (
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
}

func NewUserService() UserService {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("custom_tag", custom_tag) // 注册自定义tag
	}
	return UserService{}
}

//自定义tag函数
var custom_tag validator.Func = func(fl validator.FieldLevel) bool {
	phone, ok := fl.Field().Interface().(string)
	if ok {
		regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
		reg := regexp.MustCompile(regular)
		return reg.MatchString(phone)
	}
	return false
}

type Create struct {
	Username string `json:"user_name" binding:"required"`                                        //必须字段
	Phone    string `json:"phone" binding:"required_without=Email,omitempty,numeric,custom_tag"` //phone和email作为联系方式二选一,不能都为空,在有值时验证格式
	Email    string `json:"email" binding:"required_without=Phone,omitempty,email"`
	Password string `json:"pass_word" binding:"required"`
	Sex      string `json:"sex" binding:"required,oneof=man woman"` //性别在男、女中二选一
}

func (userService *UserService) Create(c *gin.Context) {
	var json Create
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

type Login struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"pass_word" binding:"required"`
}

func (userService *UserService) Login(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

//bind uri
type EditUri struct {
	UserId string `uri:"user_id" binding:"required"`
}
type EditJson struct {
	Phone    string    `json:"phone" binding:"omitempty,numeric"`
	Email    string    `json:"email" binding:"omitempty,email"`
	Password string    `json:"pass_word" binding:"-"`
	Birthday time.Time `json:"birthday" binding:"omitempty,lte"`
	Sex      string    `json:"sex" binding:"omitempty,oneof=man woman"`
}

func (userService *UserService) Edit(c *gin.Context) {
	var json EditJson
	var uri EditUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

type Delete struct {
	UserId string `uri:"user_id" binding:"required"`
}

func (userService *UserService) Delete(c *gin.Context) {
	var json Delete
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

type PostArticleJson struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
type PostArticlesUri struct {
	UserId string `uri:"user_id" binding:"required"`
}

func (userService *UserService) PostArticles(c *gin.Context) {
	var json PostArticleJson
	var uri PostArticlesUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

type EditArticlesUri struct {
	ArticleId string `uri:"article_id" binding:"required"`
	UserId    string `uri:"user_id" binding:"required"`
}
type EditArticlesJson struct {
	Title   string `json:"title" binding:"required_without=Content"`
	Content string `json:"content" binding:"required_without=Title"`
}

func (userService *UserService) EditArticles(c *gin.Context) {
	var uri EditArticlesUri
	var json EditArticlesJson
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}
func (userService *UserService) DeleteArticles(c *gin.Context) {
	var uri EditArticlesUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
