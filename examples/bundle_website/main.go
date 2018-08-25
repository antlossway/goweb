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

/* to compile, go to the same directory of main.go
run `rice embed-go`, a rice-box.go file is created, which contains all the data of website
then either run `go build *.go`, which will create a "main" file, run ./main
or `go install `go install github.com/antlossway/goweb/myprject`, which will create in $GOPATH/bin/myproject
*/

func main() {
	r := mux.NewRouter()
	//router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	r.PathPrefix("/go/").Handler(http.StripPrefix("/go/", http.FileServer(rice.MustFindBox("website").HTTPBox())))
	log.Fatal(http.ListenAndServe(":8080", r))
}
