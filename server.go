// asink v0.1.1-dev
//
// (c) Ground Six
//
// @package asink
// @version 0.1.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/asink/mux"
    "github.com/asink/negroni"
)

// Starts the asink built in server with a default
// on port 3000 by default
func startServer(args []string) {
    r := mux.NewRouter()
    r.HandleFunc("/", HandleRequest).
    Methods("POST").
    Host("localhost").
    Headers("Content-Type", "application/json")

    n := negroni.New()
    n.UseHandler(r)
    n.Run(":3000")
}

// Request handler for any incoming requests
// This needs to do a lot more, not just always give success!
func HandleRequest(w http.ResponseWriter, r *http.Request) {
    b, _ := ioutil.ReadAll(r.Body)
    initAsinkWithRequest(b)
    fmt.Fprintf(w, "{\"success\" : true}")
}
