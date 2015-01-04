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
	"crypto/rsa"
	"crypto/x509"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	"encoding/pem"
	"github.com/asink/go-jose"
)

type Connection struct {
	remote 		   string
	privateKeyPath string
	privateKey     *rsa.PrivateKey
	signedBody     string
}

func (c *Connection) loadPrivateKey() {
	pemData, err := ioutil.ReadFile(c.privateKeyPath)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(pemData)
    if block == nil {
        log.Fatalf("bad key data: %s", "not PEM-encoded")
    }
    if got, want := block.Type, "RSA PRIVATE KEY"; got != want {
        log.Fatalf("unknown key type %q, want %q", got, want)
    }

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        log.Fatalf("bad private key: %s", err)
    }

    c.privateKey = priv

}

func (c *Connection) signRequest(requestBody []byte) {
	signer, err := jose.NewSigner(jose.PS512, c.privateKey)
	if err != nil {
	    panic(err)
	}

	var payload = requestBody
	object, err := signer.Sign(payload)
	if err != nil {
		panic(err)
	}

	body, err := object.CompactSerialize()
	c.signedBody = string(body)

	if err != nil {
		panic(err)
	}
}

func (c Connection) makeRequest() {
	g := strings.NewReader(c.signedBody)
	_, err := http.Post(c.remote, "text/plain", g)
	if err != nil {
		panic(err)
	}
}
