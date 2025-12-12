package api

import (
	"log"
	"net/http"
	"webproject/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/login", Login)
	r.POST("/register", Register)
	r.POST("/refresh", RefreshToken)
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{

		auth.POST("/modify_password", ModifyPassword)
		auth.GET("/ping", Ping1)
	}
	log.Fatal(r.Run(":8080"))
}
func Ping1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
