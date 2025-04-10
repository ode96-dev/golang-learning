package main

import "fmt"

/*

TODO: centrifugo
- a real-time messaging server written in Go.
- let's you push data to web, mobile, and desktop clients over websocket, sse, or http-streaming - without needing to reimplement websocket logic manually
- it is like a self-hosted pub/sub server with support for:
	- broadcasting messages
	- per-user channels
	- presence (who's online)
	- authentication
	- history & recovery
	- cluster-ready
Architecture
- Go app talks to centrifugo via http api gRPC
- clients connect to centrifugo over websocket/sse
- your backend controls which users can subscribe, send messages, etc
*/

func main() {
	fmt.Println("centrifugo")
}
