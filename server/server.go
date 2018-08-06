package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/antlossway/goweb/util"
	"github.com/gorilla/mux"
)

func main() {
	//http.HandleFunc("/", handler)
	//http.HandleFunc("/deal", handlerDeal)
	//log.Fatal(http.ListenAndServe(":8080", nil))

	//	db := database{"shoes": 50, "socks": 5} //db is map[string]int
	//	log.Fatal(http.ListenAndServe("localhost:8080", db))

	//http.Handle("/", http.FileServer(http.Dir("./static")))
	//log.Fatal(http.ListenAndServe(":8080", nil))

	r := mux.NewRouter()
	//router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	r.PathPrefix("/go/").Handler(http.StripPrefix("/go/", http.FileServer(rice.MustFindBox("static").HTTPBox())))
	log.Fatal(http.ListenAndServe(":8080", r))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

//attach ServeHTTP method to databse type, so that the instance of database type satisfies
//the http.Handler interface. The handler ranges over the map and prints the items
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	joke := util.Reverse(r.URL.Path[1:])
	myrand := util.MyRandom(len(joke))
	fmt.Fprintf(w, "Hi there, I love %s and %s: %d", r.URL.Path[1:], joke, myrand)

	/*	user, ok1 := r.URL.Query().Get("user")
		password, ok2 := r.URL.Query().Get("password")
		if !ok1 or !ok2 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			login.checkuser(user, password)
		}
	*/

}
func handlerDeal(w http.ResponseWriter, r *http.Request) {
	d := util.NewDeck() // d is type deck, which "inherit" the features of slice of string []string
	//convert []string into string
	fmt.Fprintf(w, "initial deal:\n")
	for i, entry := range d {
		fmt.Fprintf(w, "%d: %s\n", i, entry)
	}

	d.Shuffle()
	fmt.Fprintf(w, "\nafter shuffle:\n")
	for i, entry := range d {
		fmt.Fprintf(w, "%d: %s\n", i, entry)
	}

	//	s := strings.Join(d, ",")
	//	fmt.Fprintf(w, "your Desk is created: %s", s)
}
