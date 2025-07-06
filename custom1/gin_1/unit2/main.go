package main

import (
	"go_study_project/custom1/gin_1/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gte=0,lte=120"`
}

var users = map[int]CreateUserInput{}
var currentID = 1

// 捕捉 panic 的中间件（全局错误处理）
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("捕获 panic：", err)
				models.Fail(c, "服务器内部错误")
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

// GET /api/v1/users/:id
func GetUserByID(c *gin.Context) {
	//panic("sadsa")
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 0 {
		//c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		models.Fail(c, "无效的用户ID")
		return
	}
	user, exist := users[id]
	if !exist {
		//c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		models.Fail(c, "用户不存在")
		return
	}
	models.OK(c, user)
}

// POST /api/v1/users
func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		models.Fail(c, err.Error())
		return
	}
	users[currentID] = input
	response := gin.H{
		"user": input,
		"id":   currentID,
	}
	currentID++
	models.OK(c, response)
}

func main() {

	r := gin.Default()
	r.Use(Recovery())
	v1 := r.Group("/api/v1")
	{
		userGroup := v1.Group("/users")
		{
			userGroup.GET("/:id", GetUserByID)
			userGroup.POST("/", CreateUser)
		}
	}

	r.Run(":8080")

}
