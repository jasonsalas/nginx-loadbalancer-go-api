package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Adder(x, y int) int {
	return x + y
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)
	fmt.Fprintln(w, "what's going on, world?")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hi", HelloHandler).Methods("GET")

	fmt.Println("server is running & listening on port 5000...")
	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
	}
}
