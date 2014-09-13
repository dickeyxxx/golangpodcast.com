package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLTemplates("templates/*")
	r.GET("/episodes/:slug", func(c *gin.Context) {
		data := gin.H{"Episode": GetEpisode(c.Params.ByName("slug"))}
		c.HTML(200, "episode.tmpl", data)
	})
	r.GET("/", func(c *gin.Context) {
		data := gin.H{"Episodes": GetEpisodes()}
		c.HTML(200, "index.tmpl", data)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
