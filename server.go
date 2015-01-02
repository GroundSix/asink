// asink v0.1.1-dev
//
// (c) Ground Six 2015
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
	"github.com/asink/mux"
	"github.com/asink/negroni"
	"io/ioutil"
	"net/http"
	"fmt"
)

var server Server

type Auth struct {
	privateKeyPath string
	publicKeyPath  string
}

type Server struct {
	Auth
	Port               string
	AuthorizedKeysPath string
}

// Starts the asink built in server with a default
// on port 3000 by default
func (s Server) Start() {
	server = s
	r := mux.NewRouter()
	r.HandleFunc("/", HandleRequest).Methods("POST")

	n := negroni.New()
	n.UseHandler(r)
	n.Run(":" + server.Port)
}

// Request handler for any incoming requests
// This needs to do a lot more, not just always give success!
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	
	//k := NewKeys()
	//println("res", string(b))
	initAsinkWithRequest(b)
	fmt.Fprintf(w, "{\"success\" : true}")
}
