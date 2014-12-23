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
	"os"
	"crypto/rsa"
	"crypto/rand"
)

type Keys struct {
	private *rsa.PrivateKey
	public  *rsa.PublicKey
}

func NewKeys() Keys {
	return Keys{}
}

func (k *Keys) Generate() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
	  panic(err)
	}
	k.private = privateKey
	k.public  = &privateKey.PublicKey
}

func (k Keys) WriteTo(filepath string) {
	if pathExists(filepath) == false {

	}
}

func (k Keys) Exist() {

}

func pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

