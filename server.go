package main

import (
	"log"
	"net/http"
	"/Users/xqy/dev/go/web/login"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
	user, ok1 := r.URL.Query().Get("user")
	password, ok2 := r.URL.Query().Get("password")
	if !ok1 or !ok2 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		login.checkuser(user, password)
	}
	
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
