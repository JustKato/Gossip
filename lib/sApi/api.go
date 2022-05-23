package sApi

import (
	"github.com/justkato/logwatch/lib/sockets"
	"github.com/justkato/logwatch/lib/types"
)

var MessagesLog map[string][]types.LogMessage = make(map[string][]types.LogMessage)

func AddLog(channel string, l types.LogMessage) {

	// Check if we have messages on this channel
	if _, ok := MessagesLog[channel]; !ok {
		// We do not have any messages, initialize them
		MessagesLog[channel] = make([]types.LogMessage, 0)
	}

	// Check if the map is at the liimt
	if len(MessagesLog[channel]) > getLogMaximumMessages() {
		// Array Shift basically
		MessagesLog[channel] = MessagesLog[channel][1:]
	}

	// Append to the messages log channel
	MessagesLog[channel] = append(MessagesLog[channel], l)

	// Broadcast the update
	BroadcastUpdate()
}

func BroadcastUpdate() {

	// Broadcast the new update to all
	sockets.BroadcastMessage(sockets.WebSocketPacket{
		SenderID: `server`,
		Event:    "update_database",
		Content:  MessagesLog,
	})

}

// GetLogs returns a list of all of the log messages that have been stored to this point in memory
func GetLogs() map[string][]types.LogMessage {
	return MessagesLog
}
