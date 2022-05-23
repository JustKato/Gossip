package sApi

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/justkato/logwatch/lib/types"
)

func HandleRouting(router *gin.RouterGroup) {

	// POST::/api/log
	router.POST("/log", func(ctx *gin.Context) {

		ch, chExists := ctx.GetPostForm(`channel`)
		if !chExists {
			ctx.JSON(400, gin.H{
				"error": "Missing 'channel', please provide a valid string for the channel",
			})
			return
		}

		content, contentExists := ctx.GetPostForm(`content`)
		if !contentExists {
			ctx.JSON(400, gin.H{
				"error": "Missing 'content', please provide a valid string for the content of the log",
			})
			return
		}

		l := types.LogMessage{
			TimeStamp: time.Now().Unix(),
			Content:   content,
			Tags:      nil,
		}

		// Add the log
		AddLog(ch, l)

		ctx.JSON(200, gin.H{
			"status":  "success",
			"message": "Log succesfully received",
		})
	})

}
