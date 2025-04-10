package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
gorilla mux
 - name parameters /user/{name}
 - query matching
 - middleware support
 - path prefixes
 - host-based routing
 - http methods filtering
*/

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "hello, %s!", name)
	})

	http.ListenAndServe(":8080", r)
}
