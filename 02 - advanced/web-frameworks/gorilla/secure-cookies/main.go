package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

/*
secure cookies
- prevents tampering via hashing and encryption
*/

var s = securecookie.New([]byte("very-secret-hash-key"), []byte("a-lot-secret-block-key"))

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{"name": "alice"}
	encoded, err := s.Encode("session", value)
	if err != nil {
		http.Error(w, "Error encoding cookie: "+err.Error(), http.StatusInternalServerError)
		fmt.Println("Encode error:", err)
		return
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: encoded,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Secure cookie set!")
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("session", cookie.Value, &value); err == nil {
			jsonValue, _ := json.Marshal(value)
			fmt.Fprintf(w, "decoded cookie: %s\n", jsonValue)
			return
		}

		http.Error(w, "failed to decode cookie", http.StatusBadRequest)
	} else {
		http.Error(w, "cookie not found", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/set", setCookieHandler)
	http.HandleFunc("/get", getCookieHandler)

	fmt.Println("starting server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
