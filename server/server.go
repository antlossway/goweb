package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antlossway/goweb/stringutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	joke := stringutil.Reverse(r.URL.Path[1:])
	fmt.Fprintf(w, "Hi there, I love %s and %s", r.URL.Path[1:], joke)
	/*	user, ok1 := r.URL.Query().Get("user")
		password, ok2 := r.URL.Query().Get("password")
		if !ok1 or !ok2 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			login.checkuser(user, password)
		}
	*/

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
