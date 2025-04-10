package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
real time communication
- means sending and receiving data immediately - in the case of chat apps, live notifications, games, collaborative tools,
- in Go, real-time is achieved with:
	- websockets
	- server-sent event(SSE)
	- log polling (less common)
	- gRPC with streaming
	- message queues + Pub/Sub (like NATS, Redis Pub/Sub)
*/

// websockets
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("upgrade error", err)
	}
	defer conn.Close()

	for {
		// read message
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read error", err)
			break
		}
		fmt.Printf("received %s\n", msg)

		// echo message back
		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println("write error", err)
			break
		}
	}
}

func main() {
	// websockets
	http.HandleFunc("/ws", handleWS)
	fmt.Println("server started at :8080")
	http.ListenAndServe(":8080", nil)

}
