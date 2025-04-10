package main

import (
	"fmt"
)

/*
Sessions
- supports memory, cookie, filesystem, and redis-based sessions
*/
// example
// var store = sessions.NewCookieStore([]byte("secret-key"))

// func sessionHandler(w http.ResponseWriter, r *http.Request) {
// 	session, _ := store.Get(r, "mysession")
// 	session.Values["authenticated"] = true
// 	session.Save(r, w)
// }

func main() {
	fmt.Println("sessions")
}
