package middleware

import (
	"errors"
	"net/http"
	"strings"
	"webproject/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

/*
	func Example1() gin.HandlerFunc {
		return func(c *gin.Context) {
			fmt.Println("example1:请求进入处理。。。")
			path := c.Request.URL.Path
			method := c.Request.Method
			fmt.Printf("正在处理：%s %s\n", method, path)
			c.Next()
			fmt.Println("example1:over")
		}
	}

	func Example2() gin.HandlerFunc {
		return func(c *gin.Context) {
			startTime := time.Now()
			fmt.Println("example2:start ticker")
			c.Next()
			latencyTime := time.Since(startTime)
			fmt.Printf("example2:请求处理耗时%v\n", latencyTime)

		}
	}
*/
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, "", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header format",
			})
			c.Abort()
			return
		}
		tokenString := parts[1]

		token, err := utils.ValidateToken(tokenString)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "token is expired",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}
		claims, err := utils.ExtractClaims(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "extract claims failed",
			})
			c.Abort()
			return
		}
		c.Set("username", claims["username"])
		c.Next()
	}

}
