package main

import (
	"fmt"
)

/*
csrf-protection
- injects csrf tokens into forms and validate them
*/

func main() {
	// CSRF := csrf.Protect([]byte("32-byte-long-auth-key"))

	// r := mux.NewRouter()
	// r.HandleFunc("/submit", submitForm)
	// http.ListenAndServe(":8000", CSRF(r))

	fmt.Println("csrf-protection")
}
