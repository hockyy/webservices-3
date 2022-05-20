package middleware

import (
	"crypto/tls"
	"fmt"
	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Log to logit.io
func LoggerToLogit() gin.HandlerFunc {
	fmt.Println("Creating new logrus instance")
	log := logrus.New()
	fmt.Println("Dialing logit")
	conn, err := tls.Dial("tcp", "47ff3430-45c7-4ef7-b3f9-f434b3e53e07-ls.logit.io:13272", &tls.Config{RootCAs: nil})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set up stash")
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "tugas-3-law"}))
	log.Hooks.Add(hook)
	return func(c *gin.Context) {
		//message := c.Param("description")
		ctx := log.WithFields(logrus.Fields{
			"method": "main",
		})
		//ctx.Info(fmt.Sprintf("[hocky.id] %s", message))
		c.Set("logit", ctx)
		c.Next()
	}
}
