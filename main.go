package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//Inisialisasi Gin Router
	router := gin.Default()

	//Middleware: Logger
	router.Use(gin.Logger())

	//Middleware: Recovery
	router.Use(gin.Recovery())

	//Route definition
	router.GET("/hello", func(c *gin.Context) {
		// 2xx -> success
		// 3xx -> redirect
		// 4xx -> bad req
		// 5xx -> internal server error
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Halo" + name + "!",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid Request Body",
			})
			return
		}
		if loginData.Email == "ardijuniawan@gmail.com" && loginData.Password == "123" {
			c.JSON(200, gin.H{
				"message": "Login Success",
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Email & Password Salah",
			})
		}
	})

	router.GET("/user", func(c *gin.Context) {
		name := c.Query("name")

		if name == "" {
			c.JSON(400, gin.H{
				"error": "Nama Parameter Kosong",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Halo, " + name + "!",
		})
	})

	//Jalankan server
	router.Run(":8080")
}
