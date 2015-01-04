// asink v0.1.1-dev
//
// (c) Ground Six 2015
//
// @package asink
// @version 0.1.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: GNU GPL v2.0
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	"github.com/asink/mux"
	"github.com/asink/negroni"
	"github.com/asink/go-jose"
	"encoding/pem"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"fmt"
	"log"
)

var server Server

type Server struct {
	Port              string
	AuthorizedKeyPath string
	RequestBody       string
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
	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}
	
	if server.AuthorizedKeyPath != "" {
		server.RequestBody = string(b)
		b = server.verifyRequest()
	}
	initAsinkWithRequest(b)
	fmt.Fprintf(w, "{\"success\" : true}")
}

// Verifies the signed request and returns the
// tasks configuration as a string
func (s Server) verifyRequest() []byte {
	pemData, err := ioutil.ReadFile(s.AuthorizedKeyPath)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(pemData)

    if block == nil {
        log.Fatalf("bad key data: %s", "not PEM-encoded")
    }
    if got, want := block.Type, "RSA PUBLIC KEY"; got != want {
        log.Fatalf("unknown key type %q, want %q", got, want)
    }

    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        log.Fatalf("bad public key: %s", err)
    }

	verifier, err := jose.NewVerifier(pub)
	if err != nil {
		panic(err)
	}

	object, err := jose.ParseSigned(s.RequestBody)
	if err != nil {
		panic(err)
	}

	output, err := verifier.Verify(object)
	if err != nil {
		panic(err)
	}

	return output
}
