package main

import (
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

/*  get some packge first
go get github.com/GeertJohan/go.rice
go get github.com/GeertJohan/go.rice/rice
go get github.com/gorilla/mux (create API)
*/

func main() {
	r := mux.NewRouter()
	//router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	r.PathPrefix("/go/").Handler(http.StripPrefix("/go/", http.FileServer(rice.MustFindBox("website").HTTPBox())))
	log.Fatal(http.ListenAndServe(":8080", r))
}
