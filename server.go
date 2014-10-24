package main

import (
	"io/ioutil"
	"net/http"
	"github.com/asink/mux"
	"github.com/asink/negroni"
)

func startServer(args []string) {
	r := mux.NewRouter()
	r.HandleFunc("/", HandleRequest)

	n := negroni.New()
	n.UseHandler(r)
	n.Run(":3000")
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	initAsinkWithRequest(b)
}
