package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
web socket
- still one of the best libraries in the Go system
*/
//example
var upgrader = websocket.Upgrader{}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	for {
		msgType, msg, _ := conn.ReadMessage()
		conn.WriteMessage(msgType, msg)
	}
}

func main() {
	fmt.Println("websocket")
}
