package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"tugas-3/middleware"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Use(middleware.LoggerToLogit())

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	r.GET("/html/:description", func(c *gin.Context) {
		tmp, _ := c.Get("logit")
		logger := tmp.(*logrus.Entry)
		//logger.Info("OK")

		message := c.Param("description")
		logger.Info(fmt.Sprintf("[Menerima pesan HTML] %s", message))
		c.HTML(http.StatusOK, "index.html", gin.H{"isi": message})
	})

	r.GET("/json/:description", func(c *gin.Context) {
		tmp, _ := c.Get("logit")
		logger := tmp.(*logrus.Entry)
		//logger.Info("OK")

		message := c.Param("description")
		logger.Info(fmt.Sprintf("[Menerima pesan JSON] %s", message))
		c.JSON(http.StatusOK, gin.H{"isi pesan": message})
	})
	r.Run()
}
