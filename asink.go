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
	//"./asink"
)

func main() {
	// Creates the root and sub commands defined
	// in options.go using cobra
	createRootCommand()
}

func initAsinkWithFile(args []string) {
	contents, err := ioutil.ReadFile(args[0])
	if (err != nil) {
		panic(err)
	}
	fmt.Println(contents)
}
