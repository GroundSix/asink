// asink v0.0.3-dev
//
// (c) Ground Six
//
// @package asink
// @version 0.0.3-dev
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
	"net/http"
	"io/ioutil"
	"./vendor/mux"
)

// Starts a very basic http server to
// accept JSON input instead of a
// static configuration file
func startServer(port string) {
	fmt.Println("Starting Asink server on port", port)
    r := mux.NewRouter()
    r.HandleFunc("/", fetchJsonBody)
    http.Handle("/", r)
    http.ListenAndServe(port, nil)
}

// Fetches the body sent in the http
// request and returns it as a string
func fetchJsonBody(w http.ResponseWriter, r *http.Request) {
	request_body, _ := ioutil.ReadAll(r.Body)
	initAsinkWithString(string(request_body))
}
