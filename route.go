package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLTemplates("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{})
	})
	r.Run(":8080")
}
