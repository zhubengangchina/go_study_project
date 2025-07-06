package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 是一个记录每个请求耗时的中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()           // 请求开始时间
		c.Next()                      // 执行后续的 Handler
		duration := time.Since(start) // 计算耗时
		fmt.Printf("请求 [%s] %s 耗时：%v\n", c.Request.Method, c.Request.URL.Path, duration)
	}
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = make(map[int]User)
var nextID = 1

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	// 1创建一个默认的路由引擎（包含日志和恢复中间件）
	r := gin.Default()

	//使用中间件
	r.Use(LoggerMiddleware())

	//注册get 路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//用户的增删改查
	//2路由注册与请求处理
	r.GET("/users", listUsers) // 查询

	r.GET("/users/:id", getUser)

	r.POST("/users", createUser)

	r.PUT("/users/:id", updateUser)

	r.DELETE("/users/:id", deleteUser)

	//3请求参数获取（Path、Query、Body、Form）
	//获取 Path 参数
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"path_id": id})
	})

	/*
		GET /users/10
		返回：{ "path_id": "10" }

	*/

	//获取 Query 参数
	r.GET("/search", func(c *gin.Context) {
		name := c.DefaultQuery("name", "guest") // 没传默认值为 guest
		age := c.Query("age")                   // 不存在则为空字符串
		c.JSON(200, gin.H{"name": name, "age": age})
	})
	/*
		GET /search?name=Tom&age=22
	*/

	//获取 JSON Body 数据（常用于 POST）
	r.POST("/login", func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"user": req.Username})
	})
	/*
		{
		  "username": "admin",
		  "password": "123456"
		}
	*/

	//获取表单参数
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		age := c.DefaultPostForm("age", "18")
		c.JSON(200, gin.H{"user": username, "age": age})
	})
	/*
		Content-Type: application/x-www-form-urlencoded
		传值为键值对形式
	*/

	//4响应输出与中间件

	//5 路由分组
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", listUsers)
		userGroup.POST("/", createUser)
		userGroup.GET("/:id", getUser)
	}
	//绑定中间件
	r.Group("/admin", LoggerMiddleware())

	//启动服务
	r.Run(":8080")
}

func listUsers(c *gin.Context) {
	list := []User{}
	for _, u := range users {
		list = append(list, u)
	}
	c.JSON(http.StatusOK, list)

}

func getUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if v, ok := users[id]; ok {
		c.JSON(http.StatusOK, v)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
	}
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = nextID
	nextID++
	users[user.ID] = user
	c.JSON(http.StatusCreated, user)
}

func updateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if _, ok := users[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = id
	users[id] = user
	c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if _, ok := users[id]; ok {
		delete(users, id)
		c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
	}
}
