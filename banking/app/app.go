package app

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {

	mux := http.NewServeMux()

	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!!")
	})
	mux.HandleFunc("/customers", GetCustomers)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
