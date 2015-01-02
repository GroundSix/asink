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
	"github.com/asink/go-homedir"
	"os"
)

// Returns the current working directory
// as a string
func getWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

// Returns the current user's home directory
// as a string
func getHomeDirectory() string {
	hd, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return hd
}
