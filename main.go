package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the API blehhhblehhh"))
}

func main() {
	r := mux.NewRouter()

	// routes
	r.HandleFunc("/", homeHandler).Methods("GET")

	// Start the server
	port := "8080"
	fmt.Println("Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
