package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/:description", func(c *gin.Context) {
		print(c.Param("description"))
		c.HTML(http.StatusOK, "index.html", gin.H{"isi": c.Param("description")})
	})

	r.GET("/json/:description", func(c *gin.Context) {
		print(c.Param("description"))
		c.JSON(http.StatusOK, gin.H{"isi": c.Param("description")})
	})
	r.Run()
}
