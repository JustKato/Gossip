package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/justkato/logwatch/lib/sockets"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("views/**/*.html")

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/ws", func(c *gin.Context) {
		sockets.WebsocketHandler(c.Writer, c.Request)
	})

	go TestTask()

	r.Run("localhost:8080")
}

func TestTask() {

	ticker := time.NewTicker(1 * time.Millisecond)

	for range ticker.C {
		sockets.BroadcastMessage(fmt.Sprintf("Current Time: %s", time.Now().Local().String()))
	}
}
