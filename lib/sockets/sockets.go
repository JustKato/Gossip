package sockets

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var SocketCache map[string]*WebSocketConnection = make(map[string]*WebSocketConnection)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func BroadcastMessage(msg string) {
	for _, element := range SocketCache {
		element.SendMessage(websocket.TextMessage, msg)
	}
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the connection to a websocket connection
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}

	// Initialize the WebSocketConnection struct
	s := WebSocketConnection{
		Socket: conn,
	}

	// Trigger the connected logic
	s.OnConnect()

	// Keep listening to the connection
	for {
		// Read the messages from the connection
		t, msg, err := conn.ReadMessage()
		if err != nil {
			// Something went wrong, most probably disconnected, kill the websocket loop.

			// Check if the entry even was cached
			if _, ok := SocketCache[conn.LocalAddr().String()]; !ok {
				// Trigger the disconnect login
				s.OnDisconnect()
			}
			break
		}

		// Trigger Read Message event
		s.OnMessage(t, msg)
	}
}

type WebSocketConnection struct {
	Socket *websocket.Conn
}

func (w *WebSocketConnection) OnConnect() {
	// Update the socket in the cache
	SocketCache[w.Socket.LocalAddr().String()] = w
}

func (w *WebSocketConnection) OnDisconnect() {
	// Delete the old reference
	delete(SocketCache, w.Socket.LocalAddr().String())

}

func (w *WebSocketConnection) OnMessage(messageType int, msg []byte) {

	// Parse the received data as a string
	finalMessage := string(msg)

	fmt.Println(finalMessage)

}

func (w *WebSocketConnection) SendMessage(messageType int, msg string) {
	w.Socket.WriteMessage(messageType, []byte(msg))
}
