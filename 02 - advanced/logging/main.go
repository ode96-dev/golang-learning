package main

import "log"

/*
logging
- log, Go's standard logging package
DRAWBACKS
- no structured (json/key-value) logging
- no levels (debug/info/warn/error)
- hard to filter on analyze at scale
*/

func main() {
	log.Println("basic log message")
	log.Printf("%s has logged in", "admin")
	log.Fatal("logs and exits the program")
}
