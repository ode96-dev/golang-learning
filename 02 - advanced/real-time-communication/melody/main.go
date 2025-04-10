package main

import (
	"net/http"

	"github.com/olahol/melody"
)

/*
Melody
- minimalistic websocket framework built on top of Gorilla websocket.
- helps in:
  - connections (clients)
  - broadcasting
  - event handling (connect, disconnect, message)
  - real-time communication
*/

func main() {
	m := melody.New()

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	m.HandleConnect(func(s *melody.Session) {
		println("new client connected")
	})

	m.HandleDisconnect(func(s *melody.Session) {
		println("client disconnected")
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		// broadcast message to everyone
		m.Broadcast(msg)
	})

	println("server started at :5500")
	http.ListenAndServe(":5500", nil)

}
