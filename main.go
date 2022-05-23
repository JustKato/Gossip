package main

import (
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/justkato/logwatch/lib/sApi"
	"github.com/justkato/logwatch/lib/sockets"
)

var doUpdate int64 = math.MaxInt64

func main() {

	// Initialize the Default GIN Configuration
	r := gin.Default()

	// Load in the HTML Templates
	r.LoadHTMLGlob("views/**/*.html")

	// Serve static Files
	r.Static("/static", "./static")

	// Handle homepage landing
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Handle WebSockets
	r.GET("/ws", func(c *gin.Context) {
		sockets.WebsocketHandler(c.Writer, c.Request)

		go func() {
			time.Sleep(time.Second)
			sApi.BroadcastUpdate()
		}()

	})

	// Route the API calls
	sApi.HandleRouting(r.Group("/api"))

	// Start running the program
	r.Run("localhost:8080")
}
