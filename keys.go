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
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
)

type Keys struct {
	path    string
	private *rsa.PrivateKey
	public  *rsa.PublicKey
}

// Creates a new instance of keys with a given
// path as it's only param
func newKeys(path string) Keys {
	k := Keys{}
	k.path = path
	return k
}

// Generates a public/private key pair for
// the Keys object
func (k *Keys) generate() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	k.private = privateKey
	k.public = &privateKey.PublicKey
}

// Writes public key to file system
func (k Keys) writePublicKey() {
	publicAsn, err := x509.MarshalPKIXPublicKey(k.public)
	if err != nil {
		panic(err)
	} else {
		publicPem := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicAsn,
		})
		ioutil.WriteFile(k.path+"/id_rsa.pub", publicPem, 0600)
	}
}

// Writes private key to file system
func (k Keys) writePrivateKey() {
	privatePem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(k.private),
		},
	)
	ioutil.WriteFile(k.path+"/id_rsa", privatePem, 0600)
}

// Checks to see if the keys exist or not
func (k Keys) exists() bool {
	if _, err := os.Stat(k.path); os.IsNotExist(err) {
		os.Mkdir(k.path, 0777)
	}
	if _, err := os.Stat(k.path + "/id_rsa"); os.IsNotExist(err) {
		k.writePrivateKey()
	}
	if _, err := os.Stat(k.path + "/id_rsa.pub"); os.IsNotExist(err) {
		k.writePublicKey()
	}
	return true
}
