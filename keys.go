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
	path    string
	private *rsa.PrivateKey
	public  *rsa.PublicKey
}

func newKeys(path string) Keys {
	k := Keys{}
	k.path = path
	return k
}

func (k *Keys) generate() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
	  panic(err)
	}
	k.private = privateKey
	k.public  = &privateKey.PublicKey
}

func (k Keys) writeTo(filepath string) {

}

func (k Keys) exists() bool {
	if _, err := os.Stat(k.path); os.IsNotExist(err) {
		return false
	}
	return true
}


