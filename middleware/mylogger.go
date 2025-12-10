package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

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
		fmt.Println("example2:start tiker")
		c.Next()
		latencyTime := time.Since(startTime)
		fmt.Printf("example2:请求处理耗时%v\n", latencyTime)

	}
}
