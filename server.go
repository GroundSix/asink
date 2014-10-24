package main

import (
	"fmt"
	"net/http"
	"github.com/asink/mux"
	"github.com/asink/negroni"
)

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandleRequest)

	n := negroni.New()
	n.UseHandler(r)
	fmt.Println("Starting server on port 3000")
	n.Run(":3000")
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("lol")
}
